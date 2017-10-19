package logging

import (
	"testing"
	"bytes"
)

func TestDebug(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Debug("name=", "elvis");

	if actual := buff.String(); actual != "DEBUG m=TestDebug name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestDebugf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Debugf("name=%v", "elvis");

	if actual := buff.String(); actual != "DEBUG m=TestDebugf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestInfo(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Info("name=", "elvis");

	if actual := buff.String(); actual != "INFO m=TestInfo name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestInfof(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Infof("name=%v", "elvis");

	if actual := buff.String(); actual != "INFO m=TestInfof name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestWarn(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Warning("name=", "elvis");

	if actual := buff.String(); actual != "WARNING m=TestWarn name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestWarnf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Warningf("name=%v", "elvis");

	if actual := buff.String(); actual != "WARNING m=TestWarnf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestError(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Error("name=", "elvis");

	if actual := buff.String(); actual != "ERROR m=TestError name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestErrorf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Errorf("name=%v", "elvis");

	if actual := buff.String(); actual != "ERROR m=TestErrorf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func NewNoFlagInstance(buff *bytes.Buffer) Log {
	return New(NewGologPrinter(buff, "", 0));
}

//
// static methods test
//
func TestStaticDebug(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Debug("name=", "elvis");

	expected := "DEBUG m=TestStaticDebug name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%q', actual='%q'", expected, actual)
	}
}

func TestStaticDebugf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Debugf("name=%v", "elvis");

	expected := "DEBUG m=TestStaticDebugf name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected=%s, actual=%s", expected, actual)
	}
}

func TestStaticInfo(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Info("name=", "elvis");

	expected := "INFO m=TestStaticInfo name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%s', actual='%s'", expected, actual)
	}
}

func TestStaticInfof(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Infof("name=%v", "elvis");

	expected := "INFO m=TestStaticInfof name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%s', actual='%s'", expected, actual)
	}
}

func TestStaticWarn(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Warning("name=", "elvis");

	expected := "WARNING m=TestStaticWarn name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expcted='%s', actual='%s'", expected, actual)
	}
}

func TestStaticWarnf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Warningf("name=%v", "elvis");

	expected := "WARNING m=TestStaticWarnf name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%s', actual='%s'", expected, actual)
	}
}

func TestStaticError(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Error("name=", "elvis");

	expected := "ERROR m=TestStaticError name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%s', actual='%s'", expected, actual)
	}
}

func TestStaticErrorf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Errorf("name=%v", "elvis");

	expected := "ERROR m=TestStaticErrorf name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%s', actual='%s'", buff.String(), actual)
	}
}

func GetStaticLoggerAndDisableTimeLogging(buff *bytes.Buffer) Log {
	logger := GetLog().(*nativeLogger)
	printer := logger.writer.(*gologPrinter)
	printer.SetOutput(buff)
	return logger
}