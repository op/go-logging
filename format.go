// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
	"time"
)

// TODO see Formatter interface in fmt/print.go
// TODO try text/template, maybe it have enough performance
// TODO other template systems?
// TODO make it possible to specify formats per backend?
type fmtVerb int

const (
	fmtVerbTime fmtVerb = iota
	fmtVerbLevel
	fmtVerbId
	fmtVerbPid
	fmtVerbProgram
	fmtVerbModule
	fmtVerbMessage
	fmtVerbLongfile
	fmtVerbShortfile
	fmtVerbLevelColor

	// Keep last, there are no match for these below.
	fmtVerbUnknown
	fmtVerbStatic
)

var fmtVerbs = []string{
	"time",
	"level",
	"id",
	"pid",
	"program",
	"module",
	"message",
	"longfile",
	"shortfile",
	"color",
}

const rfc3339Milli = "2006-01-02T15:04:05.999Z07:00"

var defaultVerbsLayout = []string{
	rfc3339Milli,
	"s",
	"d",
	"d",
	"s",
	"s",
	"s",
	"s",
	"s",
	"",
}

var (
	pid     = os.Getpid()
	program = filepath.Base(os.Args[0])
)

func getFmtVerbByName(name string) fmtVerb {
	for i, verb := range fmtVerbs {
		if name == verb {
			return fmtVerb(i)
		}
	}
	return fmtVerbUnknown
}

// Formatter is the required interface for a custom log record formatter.
type Formatter interface {
	Format(calldepth int, r *Record, w io.Writer) error
}

// formatter is used by all backends unless otherwise overriden.
var formatter struct {
	sync.RWMutex
	def Formatter
}

func getFormatter() Formatter {
	formatter.RLock()
	defer formatter.RUnlock()
	return formatter.def
}

var (
	// DefaultFormatter is the default formatter used and is only the message.
	DefaultFormatter Formatter = MustStringFormatter("%{message}")

	// Glog format
	GlogFormatter Formatter = MustStringFormatter("%{level:.1s}%{time:0102 15:04:05.999999} %{pid} %{shortfile}] %{message}")
)

// SetFormatter sets the default formatter for all new backends. A backend will
// fetch this value once it is needed to format a record. Note that backends
// will cache the formatter after the first point. For now, make sure to set
// the formatter before logging.
func SetFormatter(f Formatter) {
	formatter.Lock()
	defer formatter.Unlock()
	formatter.def = f
}

var formatRe *regexp.Regexp = regexp.MustCompile(`%{([a-z]+)(?::(.*?[^\\]))?}`)

type part struct {
	verb   fmtVerb
	layout string
}

// stringFormatter contains a list of parts which explains how to build the
// formatted string passed on to the logging backend.
type stringFormatter struct {
	parts []part
}

// NewStringFormatter returns a new Formatter which outputs the log record as a
// string based on the 'verbs' specified in the format string.
//
// The verbs:
//
// General:
//     %{id}        Sequence number for log message (uint64).
//     %{pid}       Process id (int)
//     %{time}      Time when log occurred (time.Time)
//     %{level}     Log level (Level)
//     %{module}    Module (string)
//     %{program}   Basename of os.Args[0] (string)
//     %{message}   Message (string)
//     %{longfile}  Full file name and line number: /a/b/c/d.go:23
//     %{shortfile} Final file name element and line number: d.go:23
//     %{color}     ANSI color based on log level
//
// For normal types, the output can be customized by using the 'verbs' defined
// in the fmt package, eg. '%{id:04d}' to make the id output be '%04d' as the
// format string.
//
// For time.Time, use the same layout as time.Format to change the time format
// when output, eg "2006-01-02T15:04:05.999Z-07:00".
//
// For the 'color' verb, the output can be adjusted to either use bold colors,
// i.e., '%{color:bold}' or to reset the ANSI attributes, i.e.,
// '%{color:reset}' Note that if you use the color verb explicitly, be sure to
// reset it or else the color state will persist past your log message.  e.g.,
// "%{color:bold}%{time:15:04:05} %{level:-8s}%{color:reset} %{message}" will
// just colorize the time and level, leaving the message uncolored

func NewStringFormatter(format string) (*stringFormatter, error) {
	var fmter = &stringFormatter{}

	// Find the boundaries of all %{vars}
	matches := formatRe.FindAllStringSubmatchIndex(format, -1)
	if matches == nil {
		return nil, errors.New("logger: invalid log format: " + format)
	}

	// Collect all variables and static text for the format
	prev := 0
	for _, m := range matches {
		start, end := m[0], m[1]
		if start > prev {
			fmter.add(fmtVerbStatic, format[prev:start])
		}

		name := format[m[2]:m[3]]
		verb := getFmtVerbByName(name)
		if verb == fmtVerbUnknown {
			return nil, errors.New("logger: unknown variable: " + name)
		}

		// Handle layout customizations or use the default. If this is not for the
		// time or color formatting, we need to prefix with %.
		layout := defaultVerbsLayout[verb]
		if m[4] != -1 {
			layout = format[m[4]:m[5]]
		}
		if verb != fmtVerbTime && verb != fmtVerbLevelColor {
			layout = "%" + layout
		}

		fmter.add(verb, layout)
		prev = end
	}
	end := format[prev:]
	if end != "" {
		fmter.add(fmtVerbStatic, end)
	}

	// Make a test run to make sure we can format it correctly.
	t, err := time.Parse(time.RFC3339, "2010-02-04T21:00:57-08:00")
	if err != nil {
		panic(err)
	}
	r := &Record{
		Id:     12345,
		Time:   t,
		Module: "logger",
		fmt:    "hello %s",
		args:   []interface{}{"go"},
	}
	if err := fmter.Format(0, r, &bytes.Buffer{}); err != nil {
		return nil, err
	}

	formatter.def = fmter

	return fmter, nil
}

// MustStringFormatter is equivalent to NewStringFormatter with a call to panic
// on error.
func MustStringFormatter(format string) *stringFormatter {
	f, err := NewStringFormatter(format)
	if err != nil {
		panic("Failed to initialized string formatter: " + err.Error())
	}
	return f
}

func (f *stringFormatter) add(verb fmtVerb, layout string) {
	f.parts = append(f.parts, part{verb, layout})
}

func (f *stringFormatter) Format(calldepth int, r *Record, output io.Writer) error {
	for _, part := range f.parts {
		if part.verb == fmtVerbStatic {
			output.Write([]byte(part.layout))
		} else if part.verb == fmtVerbTime {
			output.Write([]byte(r.Time.Format(part.layout)))
		} else if part.verb == fmtVerbLevelColor {
			if part.layout == "bold" {
				output.Write([]byte(boldcolors[r.Level]))
			} else if part.layout == "reset" {
				output.Write([]byte("\033[0m"))
			} else {
				output.Write([]byte(colors[r.Level]))
			}
		} else {
			var v interface{}
			switch part.verb {
			case fmtVerbLevel:
				v = r.Level
				break
			case fmtVerbId:
				v = r.Id
				break
			case fmtVerbPid:
				v = pid
				break
			case fmtVerbProgram:
				v = program
				break
			case fmtVerbModule:
				v = r.Module
				break
			case fmtVerbMessage:
				v = r.Message()
				break
			case fmtVerbLongfile, fmtVerbShortfile:
				_, file, line, ok := runtime.Caller(calldepth + 1)
				if !ok {
					file = "???"
					line = 0
				} else if part.verb == fmtVerbShortfile {
					file = filepath.Base(file)
				}
				v = fmt.Sprintf("%s:%d", file, line)
			default:
				panic("unhandled format part")
			}
			fmt.Fprintf(output, part.layout, v)
		}
	}
	return nil
}
