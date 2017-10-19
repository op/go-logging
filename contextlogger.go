package logging

import (
	"bytes"
	"log"
	"io"
)

type Printer interface {
	Printf(format string, args ...interface{})
	Print(args ...interface{})
}

type Log interface {

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

}

type nativeLogger struct {
	writer Printer
}

type logPrinter struct {
	log *log.Logger
}

func New(out io.Writer, prefix string, flag int) Log {
	l := &nativeLogger{&logPrinter{log:log.New(out, prefix, flag)}}
		return l;
}

func(p *logPrinter) Printf(format string, args ...interface{}) {
	p.log.Printf(format, args...)
}

func(p *logPrinter) Print(args ...interface{}) {
	p.log.Print(args...)
}

func (l *nativeLogger) Debug(args ...interface{}) {
	args = append([]interface{}{withCallerMethod(withLevel(new(bytes.Buffer), "DEBUG")).String()}, args...)
	l.writer.Print(args...)
}

func (l *nativeLogger) Debugf(format string, args ...interface{}) {
	l.writer.Printf(withFormat(withCallerMethod(withLevel(new(bytes.Buffer), "DEBUG")), format).String(), args...)
}

func (l *nativeLogger) Info(args ...interface{}) {
	args = append([]interface{}{withCallerMethod(withLevel(new(bytes.Buffer), "INFO")).String()}, args...)
	l.writer.Print(args...)
}
func (l *nativeLogger) Infof(format string, args ...interface{}) {
	l.writer.Printf(withFormat(withCallerMethod(withLevel(new(bytes.Buffer), "INFO")), format).String(), args...)
}

func (l *nativeLogger) Warning(args ...interface{}) {
	args = append([]interface{}{withCallerMethod(withLevel(new(bytes.Buffer), "WARNING")).String()}, args...)
	l.writer.Print(args...)
}
func (l *nativeLogger) Warningf(format string, args ...interface{}) {
	l.writer.Printf(withFormat(withCallerMethod(withLevel(new(bytes.Buffer), "WARNING")), format).String(), args...)
}

func (l *nativeLogger) Error(args ...interface{}) {
	args = append([]interface{}{withCallerMethod(withLevel(new(bytes.Buffer), "ERROR")).String()}, args...)
	l.writer.Print(args...)
}
func (l *nativeLogger) Errorf(format string, args ...interface{}) {
	l.writer.Printf(withFormat(withCallerMethod(withLevel(new(bytes.Buffer), "ERROR")), format).String(), args...)
}

const level = 2

// add method caller name to message
func withCallerMethod(buff *bytes.Buffer) *bytes.Buffer {
	buff.WriteString("m=")
	buff.WriteString(GetCallerFunctionNameSkippingAnnonymous(level))
	buff.WriteString(" ")
	return buff;
}

// adding level to message
func withLevel(buff *bytes.Buffer, lvl string) *bytes.Buffer {
	buff.WriteString(lvl)
	buff.WriteString(" ")
	return buff;
}

// adding format string to message
func withFormat(buff *bytes.Buffer, format string) *bytes.Buffer {
	buff.WriteString(format)
	return buff
}