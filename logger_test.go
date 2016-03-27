// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import "testing"

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
	if privateBackend.size != 1 {
		t.Errorf("privateBackend must contain something, size of backend: %d", privateBackend.size)
	}
	if "to private backend" != MemoryRecordN(privateBackend, 0).Formatted(0) {
		t.Error("must be logged to privateBackend:", MemoryRecordN(privateBackend, 0))
	}
}

func testConcurrent_Log(i int, sync *syncTestConcurrent, lvlBackend *LeveledBackend, log *Logger) {
	sync.start.Done()
	sync.start.Wait()
	for j := 0; j < 1000; j++ {
		log.SetBackend(*lvlBackend)
		log.Debug("to private backend")
	}
	sync.end.Done()
}

func TestPrivateBackend_Concurency(t *testing.T) {
	stdBackend := InitForTesting(DEBUG)
	log := MustGetLogger("test")
	privateBackend := NewMemoryBackend(10240)
	lvlBackend := AddModuleLevel(privateBackend)

	sync := &syncTestConcurrent{}
	sync.end.Add(10)
	sync.start.Add(10)
	for i := 0; i < 10; i++ {
		go testConcurrent_Log(i, sync, &lvlBackend, log)
	}
	sync.end.Wait()

	if stdBackend.size > 0 {
		t.Errorf("something in stdBackend, size of backend: %d", stdBackend.size)
	}
	if privateBackend.size != 10*1000 {
		t.Errorf("privateBackend must contain something, size of backend: %d", privateBackend.size)
	}
	if "to private backend" != MemoryRecordN(privateBackend, 0).Formatted(0) {
		t.Error("must be logged to privateBackend:", MemoryRecordN(privateBackend, 0))
	}
}
