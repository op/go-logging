package logging

//
// Who write the logs to output
//
type Printer interface {
	Printf(format string, args ...interface{})
	Print(args ...interface{})
}

//
// The package logger interface, you can create as many impl as you want
//
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