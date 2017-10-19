package logging

import (
	"bytes"
)

type nativeLogger struct {
	writer Printer
}

func New(p Printer) *nativeLogger {
	return &nativeLogger{p}
}

func (l *nativeLogger) Debug(args ...interface{}) {
	args = append([]interface{}{withCallerMethod(withLevel(new(bytes.Buffer), "DEBUG")).String()}, args...)
	l.Printer().Println(args...)
}

func (l *nativeLogger) Debugf(format string, args ...interface{}) {
	l.Printer().Printf(withFormat(withCallerMethod(withLevel(new(bytes.Buffer), "DEBUG")), format).String(), args...)
}

func (l *nativeLogger) Info(args ...interface{}) {
	args = append([]interface{}{withCallerMethod(withLevel(new(bytes.Buffer), "INFO")).String()}, args...)
	l.Printer().Println(args...)
}
func (l *nativeLogger) Infof(format string, args ...interface{}) {
	l.Printer().Printf(withFormat(withCallerMethod(withLevel(new(bytes.Buffer), "INFO")), format).String(), args...)
}

func (l *nativeLogger) Warning(args ...interface{}) {
	args = append([]interface{}{withCallerMethod(withLevel(new(bytes.Buffer), "WARNING")).String()}, args...)
	l.Printer().Println(args...)
}
func (l *nativeLogger) Warningf(format string, args ...interface{}) {
	l.Printer().Printf(withFormat(withCallerMethod(withLevel(new(bytes.Buffer), "WARNING")), format).String(), args...)
}

func (l *nativeLogger) Error(args ...interface{}) {
	args = append([]interface{}{withCallerMethod(withLevel(new(bytes.Buffer), "ERROR")).String()}, args...)
	l.Printer().Println(args...)
}
func (l *nativeLogger) Errorf(format string, args ...interface{}) {
	l.Printer().Printf(withFormat(withCallerMethod(withLevel(new(bytes.Buffer), "ERROR")), format).String(), args...)
}

func (l *nativeLogger) Printer() Printer {
	return l.writer
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
