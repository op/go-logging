// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"log"
	"os"
)

func init() {
	// TODO make a global Init() method to be less magic? or make it such that
	// if there's no backends at all configured, we could use some tricks to
	// automatically setup backends based if we have a TTY or not.
	b := SetBackend(NewLogBackend(os.Stderr, "", log.LstdFlags))
	b.SetLevel(DEBUG, "")
}
