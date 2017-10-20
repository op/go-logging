package logging

//
// Who write the logs to output
//
import (
	"os"
	"log"
	"io"
)

type Printer interface {
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	SetOutput(w io.Writer)
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

	Printer() Printer

}

var l Log = New(NewGologPrinter(os.Stdout, "", log.LstdFlags))
func Debug(args ...interface{}) {
	l.Debug(args)
}

func Debugf(format string, args ...interface{}){
	l.Debugf(format, args...)
}

func Info(args ...interface{}){
	l.Info(args...)
}

func Infof(format string, args ...interface{}){
	l.Infof(format, args...)
}

func Warning(args ...interface{}){
	l.Warning(args...)
}

func Warningf(format string, args ...interface{}) {
	l.Warningf(format, args...)
}

func Error(args ...interface{}){
	l.Error(args...)
}

func Errorf(format string, args ...interface{}){
	l.Errorf(format, args...)
}

func SetOutput(w io.Writer) {
	l.Printer().SetOutput(w)
}

//
// Change actual logger
//
func SetLog(logger Log){
	l = logger
}

//
// Returns current logs
//
func GetLog() Log {
	return l
}