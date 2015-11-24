// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package logging implements a logging infrastructure for Go. It supports
// different logging backends like syslog, file and memory. Multiple backends
// can be utilized with different log levels per backend and logger.
package logging

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

// Redactor is an interface for types that may contain sensitive information
// (like passwords), which shouldn't be printed to the log. The idea was found
// in relog as part of the vitness project.
type Redactor interface {
	Redacted() interface{}
}

// Redact returns a string of * having the same length as s.
func Redact(s string) string {
	return strings.Repeat("*", len(s))
}

var (
	// Sequence number is incremented and utilized for all log records created.
	sequenceNo uint64

	// timeNow is a customizable for testing purposes.
	timeNow = time.Now
)

// Record represents a log record and contains the timestamp when the record
// was created, an increasing id, filename and line and finally the actual
// formatted log line.
type Record struct {
	Id     uint64
	Time   time.Time
	Module string
	Level  Level

	// message is kept as a pointer to have shallow copies update this once
	// needed.
	message   *string
	args      []interface{}
	fmt       string
	formatter Formatter
	formatted string
}

// Formatted returns the string-formatted version of a record.
func (r *Record) Formatted(calldepth int) string {
	if r.formatted == "" {
		var buf bytes.Buffer
		r.formatter.Format(calldepth+1, r, &buf)
		r.formatted = buf.String()
	}
	return r.formatted
}

// Message returns a string message for outputting. Redacts any record args
// that implement the Redactor interface
func (r *Record) Message() string {
	if r.message == nil {
		// Redact the arguments that implements the Redactor interface
		for i, arg := range r.args {
			if redactor, ok := arg.(Redactor); ok == true {
				r.args[i] = redactor.Redacted()
			}
		}
		msg := fmt.Sprintf(r.fmt, r.args...)
		r.message = &msg
	}
	return *r.message
}

// Logger is a logging unit. It controls the flow of messages to a given
// (swappable) backend.
type Logger struct {
	Module      string
	backend     LeveledBackend
	haveBackend bool

	// ExtraCallDepth can be used to add additional call depth when getting the
	// calling function. This is normally used when wrapping a logger.
	ExtraCalldepth int
}

// SetBackend changes the backend of the logger.
func (l *Logger) SetBackend(backend LeveledBackend) {
	l.backend = backend
	l.haveBackend = true
}

// GetLogger creates and returns a Logger object based on the module name.
// TODO call NewLogger and remove MustGetLogger?
func GetLogger(module string) (*Logger, error) {
	return &Logger{Module: module}, nil
}

// MustGetLogger is like GetLogger but panics if the logger can't be created.
// It simplifies safe initialization of a global logger for eg. a package.
func MustGetLogger(module string) *Logger {
	logger, err := GetLogger(module)
	if err != nil {
		panic("logger: " + module + ": " + err.Error())
	}
	return logger
}

// Reset restores the internal state of the logging library.
func Reset() {
	// TODO make a global Init() method to be less magic? or make it such that
	// if there's no backends at all configured, we could use some tricks to
	// automatically setup backends based if we have a TTY or not.
	sequenceNo = 0
	b := SetBackend(NewLogBackend(os.Stderr, "", log.LstdFlags))
	b.SetLevel(DEBUG, "")
	SetFormatter(DefaultFormatter)
	timeNow = time.Now
}

// IsEnabledFor returns true if the logger is enabled for the given level.
func (l *Logger) IsEnabledFor(level Level) bool {
	return defaultBackend.IsEnabledFor(level, l.Module)
}

func (l *Logger) log(lvl Level, format string, args ...interface{}) {
	if !l.IsEnabledFor(lvl) {
		return
	}

	// Create the logging record and pass it in to the backend
	record := &Record{
		Id:     atomic.AddUint64(&sequenceNo, 1),
		Time:   timeNow(),
		Module: l.Module,
		Level:  lvl,
		fmt:    format,
		args:   args,
	}

	// TODO use channels to fan out the records to all backends?
	// TODO in case of errors, do something (tricky)

	// calldepth=2 brings the stack up to the caller of the level
	// methods, Info(), Fatal(), etc.
	// ExtraCallDepth allows this to be extended further up the stack in case we
	// are wrapping these methods, eg. to expose them package level
	if l.haveBackend {
		l.backend.Log(lvl, 2+l.ExtraCalldepth, record)
		return
	}

	defaultBackend.Log(lvl, 2+l.ExtraCalldepth, record)
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func (l *Logger) Fatal(args ...interface{}) {
	s := fmt.Sprint(args...)
	l.log(CRITICAL, "%s", s)
	os.Exit(1)
}

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log(CRITICAL, format, args...)
	os.Exit(1)
}

// Panic is equivalent to l.Critical(fmt.Sprint()) followed by a call to panic().
func (l *Logger) Panic(args ...interface{}) {
	s := fmt.Sprint(args...)
	l.log(CRITICAL, "%s", s)
	panic(s)
}

// Panicf is equivalent to l.Critical followed by a call to panic().
func (l *Logger) Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	l.log(CRITICAL, "%s", s)
	panic(s)
}

// Critical logs a message using CRITICAL as log level. (fmt.Sprint())
func (l *Logger) Critical(args ...interface{}) {
	s := fmt.Sprint(args...)
	l.log(CRITICAL, "%s", s)
}

// Criticalf logs a message using CRITICAL as log level.
func (l *Logger) Criticalf(format string, args ...interface{}) {
	l.log(CRITICAL, format, args...)
}

// Error logs a message using ERROR as log level. (fmt.Sprint())
func (l *Logger) Error(args ...interface{}) {
	s := fmt.Sprint(args...)
	l.log(ERROR, "%s", s)
}

// Errorf logs a message using ERROR as log level.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Errorf logs a message using ERROR as log level.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Warning logs a message using WARNING as log level.
func (l *Logger) Warning(args ...interface{}) {
	s := fmt.Sprint(args...)
	l.log(WARNING, "%s", s)
}

// Warningf logs a message using WARNING as log level.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.log(WARNING, format, args...)
}

// Warningf logs a message using WARNING as log level.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.log(WARNING, format, args...)
}

// Notice logs a message using NOTICE as log level.
func (l *Logger) Notice(args ...interface{}) {
	s := fmt.Sprint(args...)
	l.log(NOTICE, "%s", s)
}

// Noticef logs a message using NOTICE as log level.
func (l *Logger) Noticef(format string, args ...interface{}) {
	l.log(NOTICE, format, args...)
}

// Noticef logs a message using NOTICE as log level.
func (l *Logger) Noticef(format string, args ...interface{}) {
	l.log(NOTICE, format, args...)
}

// Info logs a message using INFO as log level.
func (l *Logger) Info(args ...interface{}) {
	s := fmt.Sprint(args...)
	l.log(INFO, "%s", s)
}

// Infof logs a message using INFO as log level.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Infof logs a message using INFO as log level.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Debug logs a message using DEBUG as log level.
func (l *Logger) Debug(args ...interface{}) {
	s := fmt.Sprint(args...)
	l.log(DEBUG, "%s", s)
}

// Debugf logs a message using DEBUG as log level.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Debugf logs a message using DEBUG as log level.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

func init() {
	Reset()
}
