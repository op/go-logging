package logging

import (
	"context"
	"sync/atomic"
	"os"
	"fmt"
	"strings"
	"bytes"
)

type ContextLogger struct {
	logger Log
	ctx context.Context
	id uint64
}

func (l *ContextLogger) Critical(args ...interface{}) {
	l.logger.Critical(getArgs(l, args...)...)
}
func (l *ContextLogger) Criticalf(format string, args ...interface{}) {
	l.logger.Criticalf(getIdConcat(l, format, -1), args...)
}
func (l *ContextLogger) Debug(args ...interface{}) {
	l.logger.Debug(getArgs(l, args...)...)
}
func (l *ContextLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(getIdConcat(l, format, -1), args...)
}
func (l *ContextLogger) Error(args ...interface{}) {
	l.logger.Error(getArgs(l, args...)...)
}
func (l *ContextLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(getIdConcat(l, format, -1), args...)
}
func (l *ContextLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(getArgs(l, args...)...)
	os.Exit(1)
}
func (l *ContextLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(getIdConcat(l, format, -1), args...)
	os.Exit(1)
}
func (l *ContextLogger) Info(args ...interface{}) {
	l.logger.Info(getArgs(l, args...)...)
}
func (l *ContextLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(getIdConcat(l, format, -1), args...)

}

func (l *ContextLogger) Notice(args ...interface{}) {
	l.logger.Notice(getArgs(l, args...)...)
}
func (l *ContextLogger) Noticef(format string, args ...interface{}) {
	l.logger.Noticef(getIdConcat(l, format, -1), args...)
}
func (l *ContextLogger) Panic(args ...interface{}) {
	l.logger.Panic(args)
	panic(fmt.Sprint(args...))
}
func (l *ContextLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(getIdConcat(l, format, -1), args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *ContextLogger) Warning(args ...interface{}) {
	l.logger.Warning(getArgs(l, args...)...)
}
func (l *ContextLogger) Warningf(format string, args ...interface{}) {
	l.logger.Warningf(getIdConcat(l, format, -1), args...)
}


func getArgs(l *ContextLogger, args ...interface{}) []interface{} {
	var ar  []interface{}
	ar = append(ar, strings.Trim(getId(l), " "))
	ar = append(ar, args...)
	return ar;
}

func getId(l *ContextLogger) string {
	return getIdConcat(l, "", 4)
}

func getIdConcat(l *ContextLogger, append string, level int ) string {
	if level == -1{
		level = 2
	}
	var buffer bytes.Buffer
	buffer.WriteString("id=")
	buffer.WriteString(fmt.Sprintf("%d, ", l.id))
	buffer.WriteString("m=")
	buffer.WriteString(GetCallerFunctionNameSkippingAnnonymous(level))
	buffer.WriteString(" ")
	buffer.WriteString(append)
	return buffer.String()
}

var logger Log;
var instanced int32 = 0
func NewLog(ctx context.Context) Log {
	if(atomic.CompareAndSwapInt32(&instanced, 0, 1)){
		logger = MustGetLogger("main")
	}
	return NewLogWithContext(ctx, logger)
}
func NewLogWithLogger(logger Log) Log {

	ctx := NewContext()
	var ctxLogger ContextLogger
	if id := ctx.Value(idKey); id != nil {
		ctxLogger = ContextLogger{logger, ctx, id.(uint64)}
	} else {
		ctxContext := ContextWithId(ctx)
		ctxLogger = ContextLogger{logger, ctxContext, ctxContext.Value(id).(uint64)}
	}
	return &ctxLogger
}

func NewLogWithContext(ctx context.Context, logger Log) Log {

	var ctxLogger ContextLogger
	if id := ctx.Value(idKey); id != nil {
		ctxLogger = ContextLogger{logger, ctx, id.(uint64)}
	} else {
		ctxContext := ContextWithId(ctx)
		ctxLogger = ContextLogger{logger, ctxContext, ctxContext.Value(id).(uint64)}
	}
	return &ctxLogger
}

const idKey = "ctx_id"
var id uint64 = 0;
func NewContext() context.Context {
	return ContextWithId(context.Background())
}

func ContextWithId(parent context.Context) context.Context {
	newId := atomic.AddUint64(&id, 1)
	ctx := context.WithValue(parent, idKey, newId)
	return ctx
}

