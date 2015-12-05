// +build windows
// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"bytes"
	"io"
	"log"
	"syscall"
)

type FileDescriptorGetter interface {
	Fd() uintptr
}

var (
	kernel32DLL                 = syscall.NewLazyDLL("kernel32.dll")
	setConsoleTextAttributeProc = kernel32DLL.NewProc("SetConsoleTextAttribute")
)

// TODO initialize here
var win_colors []WORD
var win_boldcolors []WORD


type WORD uint16

const (
	// Character attributes
	// Note:
	// -- The attributes are combined to produce various colors (e.g., Blue + Green will create Cyan).
	//    Clearing all foreground or background colors results in black; setting all creates white.
	// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms682088(v=vs.85).aspx#_win32_character_attributes.
	fgBlack     WORD = 0x0000
	fgBlue      WORD = 0x0001
	fgGreen     WORD = 0x0002
	fgCyan      WORD = 0x0003
	fgRed       WORD = 0x0004
	fgMagenta   WORD = 0x0005
	fgYellow    WORD = 0x0006
	fgWhite     WORD = 0x0007
	fgIntensity WORD = 0x0008
	fgMask      WORD = 0x000F
)

// LogBackend utilizes the standard log module.
type LogBackend struct {
	Logger *log.Logger
	Color  bool
	Handle uintptr
}

// NewLogBackend creates a new LogBackend.
func NewLogBackend(out io.Writer, prefix string, flag int) *LogBackend {
	var handle uintptr
	if fdg, ok := interface{}(out).(FileDescriptorGetter); ok {
		handle = fdg.Fd()
	}
	return &LogBackend{Logger: log.New(out, prefix, flag), Handle: handle}
}

func (b *LogBackend) Log(level Level, calldepth int, rec *Record) error {
	if b.Color {
		buf := &bytes.Buffer{}
		setConsoleTextAttribute(b.Handle, win_colors[level])
		buf.Write([]byte(rec.Formatted(calldepth + 1)))
		// For some reason, the Go logger arbitrarily decided "2" was the correct
		// call depth...
		err := b.Logger.Output(calldepth+2, buf.String())
		setConsoleTextAttribute(b.Handle, fgWhite)
		return err
	}
	return b.Logger.Output(calldepth+2, rec.Formatted(calldepth+1))
}

func init() {
	init_colors()
	win_colors = []WORD{
		INFO:     fgWhite,
		CRITICAL: fgMagenta,
		ERROR:    fgRed,
		WARNING:  fgYellow,
		NOTICE:   fgGreen,
		DEBUG:    fgCyan,
	}
	win_boldcolors = []WORD{
		INFO:     fgWhite | fgIntensity,
		CRITICAL: fgMagenta | fgIntensity,
		ERROR:    fgRed | fgIntensity,
		WARNING:  fgYellow | fgIntensity,
		NOTICE:   fgGreen | fgIntensity,
		DEBUG:    fgCyan | fgIntensity,
	}
}

// setConsoleTextAttribute sets the attributes of characters written to the
// console screen buffer by the WriteFile or WriteConsole function.
// See http://msdn.microsoft.com/en-us/library/windows/desktop/ms686047(v=vs.85).aspx.
func setConsoleTextAttribute(handle uintptr, attribute WORD) error {
	r1, r2, err := setConsoleTextAttributeProc.Call(handle, uintptr(attribute), 0)
	use(attribute)
	return checkError(r1, r2, err)
}

// checkError evaluates the results of a Windows API call and returns the error if it failed.
func checkError(r1, r2 uintptr, err error) error {
	// Windows APIs return non-zero to indicate success
	if r1 != 0 {
		return nil
	}

	// Return the error if provided, otherwise default to EINVAL
	if err != nil {
		return err
	}
	return syscall.EINVAL
}

// use is a no-op, but the compiler cannot see that it is.
// Calling use(p) ensures that p is kept live until that point.
func use(p interface{}) {}

