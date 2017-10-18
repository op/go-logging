package logging

import (
	"bytes"
)

type nativeLogger struct {
	logger Log
}

func (l *nativeLogger) Debug(args ...interface{}) {
	l.logger.Debug(withMethod("")...)
}

func (l *nativeLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(withMethod(format), args...)
}

func (l *nativeLogger) Info(args ...interface{}) {
	l.logger.Info(withMethod("")...)
}
func (l *nativeLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(withMethod(format), args...)
}

func (l *nativeLogger) Warning(args ...interface{}) {
	l.logger.Warning(withMethod("")...)
}
func (l *nativeLogger) Warningf(format string, args ...interface{}) {
	l.logger.Warningf(withMethod(format), args...)
}

func (l *nativeLogger) Error(args ...interface{}) {
	l.logger.Error(withMethod("")...)
}
func (l *nativeLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(withMethod(format), args...)
}


const level = 2
func withMethod(append string) string {
	var buffer bytes.Buffer
	buffer.WriteString("m=")
	buffer.WriteString(GetCallerFunctionNameSkippingAnnonymous(level))
	buffer.WriteString(" ")
	buffer.WriteString(append)
	return buffer.String()
}
