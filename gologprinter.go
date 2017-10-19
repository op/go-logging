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

func(p *gologPrinter) Println(args ...interface{}) {
	p.log.Println(args...)
}

func(p *gologPrinter) SetFlags(flags int){
	p.log.SetFlags(flags)
}

func(p *gologPrinter) SetOutput(out io.Writer){
	p.log.SetOutput(out)
}

func NewGologPrinter(out io.Writer, prefix string, flag int) *gologPrinter {
	return &gologPrinter{log: log.New(out, prefix, flag)}
}
