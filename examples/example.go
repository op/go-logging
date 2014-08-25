package main

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = "%{color}%{time:15:04:05.000000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}"

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func main() {
	// Setup one stderr and one syslog backend and combine them both into one
	// logging backend. By default stderr is used with the standard log flag.
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	syslogBackend, err := logging.NewSyslogBackend("")
	if err != nil {
		log.Fatal(err)
	}
	logging.SetBackend(logBackend, syslogBackend)
	logging.SetFormatter(logging.MustStringFormatter(format))

	// Run one with debug setup for "test" and one with error.
	for _, level := range []logging.Level{logging.DEBUG, logging.ERROR} {
		logging.SetLevel(level, "example")

		log.Critical("crit")
		log.Error("err")
		log.Warning("warning")
		log.Notice("notice")
		log.Info("info")
		log.Debug("debug %s", Password("secret"))
	}
}
