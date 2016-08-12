// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"sync/atomic"
)

// defaultBackend is the backend used for all logging calls.
var defaultBackend atomic.Value

// Backend is the interface which a log backend need to implement to be able to
// be used as a logging backend.
type Backend interface {
	Log(Level, int, *Record) error
}

// SetBackend replaces the backend currently set with the given new logging
// backend.
func SetBackend(backends ...Backend) LeveledBackend {
	var backend Backend
	if len(backends) == 1 {
		backend = backends[0]
	} else {
		backend = MultiLogger(backends...)
	}

	newDefaultBackend := AddModuleLevel(backend)
	defaultBackend.Store(&newDefaultBackend)

	return *defaultBackend.Load().(*LeveledBackend)
}

// SetLevel sets the logging level for the specified module. The module
// corresponds to the string specified in GetLogger.
func SetLevel(level Level, module string) {
	(*defaultBackend.Load().(*LeveledBackend)).SetLevel(level, module)
}

// GetLevel returns the logging level for the specified module.
func GetLevel(module string) Level {
	return (*defaultBackend.Load().(*LeveledBackend)).GetLevel(module)
}
