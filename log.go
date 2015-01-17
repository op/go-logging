// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// TODO initialize here
var colors []string
var boldcolors []string

type color int

const (
	colorBlack = (iota + 30)
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite
)

// LogBackend utilizes the standard log module.
type LogBackend struct {
	Logger *log.Logger
	Color  bool
}

// NewLogBackend creates a new LogBackend.
func NewLogBackend(out io.Writer, prefix string, flag int) *LogBackend {
	return &LogBackend{Logger: log.New(out, prefix, flag)}
}

func (b *LogBackend) Log(level Level, calldepth int, rec *Record) error {
	if b.Color {
		buf := &bytes.Buffer{}
		buf.Write([]byte(colors[level]))
		buf.Write([]byte(rec.Formatted(calldepth + 1)))
		buf.Write([]byte("\033[0m"))
		// For some reason, the Go logger arbitrarily decided "2" was the correct
		// call depth...
		return b.Logger.Output(calldepth+2, buf.String())
	} else {
		return b.Logger.Output(calldepth+2, rec.Formatted(calldepth+1))
	}
	panic("should not be reached")
}

func colorSeq(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

func colorSeqBold(color color) string {
	return fmt.Sprintf("\033[%d;1m", int(color))
}

type DefaultLogBackend struct {
	file *os.File
	path string
}

// Create a new backend that uses the default golang logger
//
// The default golang logger exposes a SetOutput method which allows us to
// effectively reopen the log file on demand, whereas a regular golang logger
// does not. Thus this allows us to use the default logger for that purpose
func NewDefaultLogBackend(path string, prefix string, flag int) (*DefaultLogBackend, error) {
	ret := &DefaultLogBackend{
		path: path,
	}

	log.SetPrefix(prefix)
	log.SetFlags(flag)

	err := ret.Reopen()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (f *DefaultLogBackend) Log(level Level, calldepth int, rec *Record) error {
	log.Print(rec.Formatted(calldepth + 1))
	return nil
}

func (f *DefaultLogBackend) Reopen() (err error) {
	var new_file *os.File

	new_file, err = os.OpenFile(f.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0640)
	if err != nil {
		return
	}

	// Switch to new output before closing
	log.SetOutput(new_file)

	if f.file != nil {
		f.file.Close()
	}

	f.file = new_file

	return nil
}

func (f *DefaultLogBackend) Close() {
	// Discard logs before closing
	log.SetOutput(ioutil.Discard)

	if f.file != nil {
		f.file.Close()
	}

	f.file = nil
}

func init() {
	colors = []string{
		CRITICAL: colorSeq(colorMagenta),
		ERROR:    colorSeq(colorRed),
		WARNING:  colorSeq(colorYellow),
		NOTICE:   colorSeq(colorGreen),
		DEBUG:    colorSeq(colorCyan),
	}
	boldcolors = []string{
		CRITICAL: colorSeqBold(colorMagenta),
		ERROR:    colorSeqBold(colorRed),
		WARNING:  colorSeqBold(colorYellow),
		NOTICE:   colorSeqBold(colorGreen),
		DEBUG:    colorSeqBold(colorCyan),
	}
}
