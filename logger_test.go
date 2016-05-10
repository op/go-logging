// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"testing"
	"time"
)

type Password string

func (p Password) Redacted() interface{} {
	return Redact(string(p))
}

func TestSequenceNoOverflow(t *testing.T) {
	// Forcefully set the next sequence number to the maximum
	backend := InitForTesting(DEBUG)
	sequenceNo = ^uint64(0)

	log := MustGetLogger("test")
	log.Debug("test")

	if MemoryRecordN(backend, 0).ID != 0 {
		t.Errorf("Unexpected sequence no: %v", MemoryRecordN(backend, 0).ID)
	}
}

func TestRedact(t *testing.T) {
	backend := InitForTesting(DEBUG)
	password := Password("123456")
	log := MustGetLogger("test")
	log.Debug("foo", password)
	if "foo ******" != MemoryRecordN(backend, 0).Formatted(0) {
		t.Errorf("redacted line: %v", MemoryRecordN(backend, 0))
	}
}

func TestRedactf(t *testing.T) {
	backend := InitForTesting(DEBUG)
	password := Password("123456")
	log := MustGetLogger("test")
	log.Debugf("foo %s", password)
	if "foo ******" != MemoryRecordN(backend, 0).Formatted(0) {
		t.Errorf("redacted line: %v", MemoryRecordN(backend, 0).Formatted(0))
	}
}

func TestPrivateBackend(t *testing.T) {
	stdBackend := InitForTesting(DEBUG)
	log := MustGetLogger("test")
	privateBackend := NewMemoryBackend(10240)
	lvlBackend := AddModuleLevel(privateBackend)
	lvlBackend.SetLevel(DEBUG, "")
	log.SetBackend(lvlBackend)
	log.Debug("to private backend")
	if stdBackend.size > 0 {
		t.Errorf("something in stdBackend, size of backend: %d", stdBackend.size)
	}
	if "to private baсkend" == MemoryRecordN(privateBackend, 0).Formatted(0) {
		t.Error("logged to defaultBackend:", MemoryRecordN(privateBackend, 0))
	}
}

func TestLogger_SetTimerNow(t *testing.T) {
	_ = InitForTesting(DEBUG)
	_ = MustGetLogger("test")

	if now := timeNow(); now != time.Unix(0, 0).UTC() {
		t.Error("test timeNow incorrect", now)
	}

	SetTimeNow(func() time.Time {
		return time.Unix(1, 1)
	})

	if now := timeNow(); now != time.Unix(1, 1) {
		t.Error("test timeNow dit not get overwritten", now)
	}
}
