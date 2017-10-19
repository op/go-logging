package logging

import (
	"log"
	"io"
)

//
// log using golang native logger solution
//
type gologPrinter struct {
	log *log.Logger
}

func(p *gologPrinter) Printf(format string, args ...interface{}) {
	p.log.Printf(format, args...)
}

func(p *gologPrinter) Print(args ...interface{}) {
	p.log.Print(args...)
}

func NewGologPrinter(out io.Writer, prefix string, flag int) *log.Logger {
	return log.New(out, prefix, flag)
}