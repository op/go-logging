// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !unix,!nacl,!plan9

package logging

import (
	"os"
)

// NewSyslogBackend (Windows version) conforms to the same signature as the Unix version, but merely returns a LogBackend that is set to standard out.
func NewSyslogBackend(prefix string) (b *LogBackend, err error) {
	return NewLogBackend(os.Stdout, prefix, 0), nil
}

// NewSyslogBackendPriority (Windows version) conforms to the same signature as the Unix version, but merely returns a LogBackend that is set to standard out. (The priority parameter is ignored.)
func NewSyslogBackendPriority(prefix string, priority interface{}) (b *LogBackend, err error) {
	return NewLogBackend(os.Stdout, prefix, 0), nil
}
