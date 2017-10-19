package logging

//
// Who write the logs to output
//
import "os"

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

var l Log = New(NewGologPrinter(os.Stdout, "", 0))
func Debug(args ...interface{}) {
	l.Debug(args)
}

func Debugf(format string, args ...interface{}){
	l.Debugf(format, args)
}

func Info(args ...interface{}){
	l.Info(args)
}

func Infof(format string, args ...interface{}){
	l.Infof(format, args)
}

func Warning(args ...interface{}){
	l.Warning( args)
}

func Warningf(format string, args ...interface{}) {
	l.Warningf(format, args)
}

func Error(args ...interface{}){
	l.Error(args)
}

func Errorf(format string, args ...interface{}){
	l.Errorf(format, args)
}

func SetLog(logger Log){
	l = logger
}

func GetLog() Log {
	return l
}