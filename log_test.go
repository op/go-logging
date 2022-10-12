// Copyright 2013, Örjan Persson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"bytes"
	"io"
	"log"
	"strings"
	"testing"
)

func TestLogCalldepth(t *testing.T) {
	buf := &bytes.Buffer{}
	SetBackend(NewLogBackend(buf, "", log.Lshortfile))
	SetFormatter(MustStringFormatter("%{shortfile} %{level} %{message}"))

	log := MustGetLogger("test")
	log.Info("test filename")

	parts := strings.SplitN(buf.String(), " ", 2)

	// Verify that the correct filename is registered by the stdlib logger
	if !strings.HasPrefix(parts[0], "log_test.go:") {
		t.Errorf("incorrect filename: %s", parts[0])
	}
	// Verify that the correct filename is registered by go-logging
	if !strings.HasPrefix(parts[1], "log_test.go:") {
		t.Errorf("incorrect filename: %s", parts[1])
	}
}

func c(log *Logger) { log.Info("test callpath") }
func b(log *Logger) { c(log) }
func a(log *Logger) { b(log) }

func rec(log *Logger, r int) {
	if r == 0 {
		a(log)
		return
	}
	rec(log, r-1)
}

func testCallpath(t *testing.T, format string, expect string) {
	buf := &bytes.Buffer{}
	SetBackend(NewLogBackend(buf, "", log.Lshortfile))
	SetFormatter(MustStringFormatter(format))

	logger := MustGetLogger("test")
	rec(logger, 6)

	parts := strings.SplitN(buf.String(), " ", 3)

	// Verify that the correct filename is registered by the stdlib logger
	if !strings.HasPrefix(parts[0], "log_test.go:") {
		t.Errorf("incorrect filename: %s", parts[0])
	}
	// Verify that the correct callpath is registered by go-logging
	if !strings.HasPrefix(parts[1], expect) {
		t.Errorf("incorrect callpath: %s missing prefix %s", parts[1], expect)
	}
	// Verify that the correct message is registered by go-logging
	if !strings.HasPrefix(parts[2], "test callpath") {
		t.Errorf("incorrect message: %s", parts[2])
	}
}

func TestLogCallpath(t *testing.T) {
	testCallpath(t, "%{callpath} %{message}", "TestLogCallpath.String.rec...a.b.c")
	testCallpath(t, "%{callpath:-1} %{message}", "TestLogCallpath.String.rec...a.b.c")
	testCallpath(t, "%{callpath:0} %{message}", "TestLogCallpath.String.rec...a.b.c")
	testCallpath(t, "%{callpath:1} %{message}", "~.c")
	testCallpath(t, "%{callpath:2} %{message}", "~.c.c")
	testCallpath(t, "%{callpath:3} %{message}", "~.b.c.c")
}

func BenchmarkLogMemoryBackendIgnored(b *testing.B) {
	backend := SetBackend(NewMemoryBackend(1024))
	backend.SetLevel(INFO, "")
	RunLogBenchmark(b)
}

func BenchmarkLogMemoryBackend(b *testing.B) {
	backend := SetBackend(NewMemoryBackend(1024))
	backend.SetLevel(DEBUG, "")
	RunLogBenchmark(b)
}

func BenchmarkLogChannelMemoryBackend(b *testing.B) {
	channelBackend := NewChannelMemoryBackend(1024)
	backend := SetBackend(channelBackend)
	backend.SetLevel(DEBUG, "")
	RunLogBenchmark(b)
	channelBackend.Flush()
}

func BenchmarkLogLeveled(b *testing.B) {
	backend := SetBackend(NewLogBackend(io.Discard, "", 0))
	backend.SetLevel(INFO, "")

	RunLogBenchmark(b)
}

func BenchmarkLogLogBackend(b *testing.B) {
	backend := SetBackend(NewLogBackend(io.Discard, "", 0))
	backend.SetLevel(DEBUG, "")
	RunLogBenchmark(b)
}

func BenchmarkLogLogBackendColor(b *testing.B) {
	colorizer := NewLogBackend(io.Discard, "", 0)
	colorizer.Color = true
	backend := SetBackend(colorizer)
	backend.SetLevel(DEBUG, "")
	RunLogBenchmark(b)
}

func BenchmarkLogLogBackendStdFlags(b *testing.B) {
	backend := SetBackend(NewLogBackend(io.Discard, "", log.LstdFlags))
	backend.SetLevel(DEBUG, "")
	RunLogBenchmark(b)
}

func BenchmarkLogLogBackendLongFileFlag(b *testing.B) {
	backend := SetBackend(NewLogBackend(io.Discard, "", log.Llongfile))
	backend.SetLevel(DEBUG, "")
	RunLogBenchmark(b)
}

func RunLogBenchmark(b *testing.B) {
	password := Password("foo")
	log := MustGetLogger("test")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("log line for %d and this is rectified: %s", i, password)
	}
}

func BenchmarkLogFixed(b *testing.B) {
	backend := SetBackend(NewLogBackend(io.Discard, "", 0))
	backend.SetLevel(DEBUG, "")

	RunLogBenchmarkFixedString(b)
}

func BenchmarkLogFixedIgnored(b *testing.B) {
	backend := SetBackend(NewLogBackend(io.Discard, "", 0))
	backend.SetLevel(INFO, "")
	RunLogBenchmarkFixedString(b)
}

func RunLogBenchmarkFixedString(b *testing.B) {
	log := MustGetLogger("test")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("some random fixed text")
	}
}
