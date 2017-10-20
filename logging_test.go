package logging

import (
	"testing"
	"bytes"
	"os"
	"log"
)

func TestDebug(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Debug("name=", "elvis");

	expected := "DEBUG m=TestDebug  name= elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, full=%s, actual=%s", expected, actual)
	}
}

func TestDebugf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Debugf("name=%v", "elvis");

	expected := "DEBUG m=TestDebugf name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, full=%s, actual=%s", expected, actual)
	}
}

func TestInfo(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Info("name=", "elvis");

	expected := "INFO m=TestInfo  name= elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, full=%s, actual=%s", expected, actual)
	}
}

func TestInfof(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Infof("name=%v", "elvis");

	expected := "INFO m=TestInfof name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, full=%s, actual=%s", expected, actual)
	}
}

func TestWarn(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Warning("name=", "elvis");

	expected := "WARNING m=TestWarn  name= elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, full=%s, actual=%s", expected, actual)
	}
}

func TestWarnf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Warningf("name=%v", "elvis");

	expected := "WARNING m=TestWarnf name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, full=%s, actual=%s", expected, actual)
	}
}

func TestError(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Error("name=", "elvis");

	expected := "ERROR m=TestError  name= elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, full=%s, actual=%s", expected, actual)
	}
}

func TestErrorf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := NewNoFlagInstance(buff)
	logger.Errorf("name=%v", "elvis");

	expected := "ERROR m=TestErrorf name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, full=%s, actual=%s", expected, actual)
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

	expected := "DEBUG m=TestStaticDebug  name= elvis\n"
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

	expected := "INFO m=TestStaticInfo  name= elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%q', actual='%q'", expected, actual)
	}
}

func TestStaticInfof(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Infof("name=%v", "elvis");

	expected := "INFO m=TestStaticInfof name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%q', actual='%q'", expected, actual)
	}
}

func TestStaticWarn(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Warning("name=", "elvis");

	expected := "WARNING m=TestStaticWarn  name= elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expcted='%q', actual='%q'", expected, actual)
	}
}

func TestStaticWarnf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Warningf("name=%v", "elvis");

	expected := "WARNING m=TestStaticWarnf name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%q', actual='%q'", expected, actual)
	}
}

func TestStaticError(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Error("name=", "elvis");

	expected := "ERROR m=TestStaticError  name= elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%q', actual='%q'", expected, actual)
	}
}

func TestStaticErrorf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := GetStaticLoggerAndDisableTimeLogging(buff)
	logger.Errorf("name=%v", "elvis");

	expected := "ERROR m=TestStaticErrorf name=elvis\n"
	if actual := buff.String(); actual != expected {
		t.Errorf("log format not expected, expected='%q', actual='%q'", expected, actual)
	}
}

func GetStaticLoggerAndDisableTimeLogging(buff *bytes.Buffer) Log {
	logger := GetLog()
	printer := logger.Printer().(*gologPrinter)
	printer.SetFlags(0)
	printer.SetOutput(buff)
	return logger
}

func ExampleDebugf() {
	printer := GetLog().Printer().(*gologPrinter)
	printer.SetOutput(os.Stdout)
	printer.SetFlags(0)

	Debugf("name=%q, age=%d", "John\nZucky", 21)

	// Output:
	// DEBUG m=Debugf name="John\nZucky", age=21
}

func BenchmarkDebugf(b *testing.B) {

	//go pprof.Lookup("block").WriteTo(os.Stdout, 2)
	//f, err := os.Open("./cpu.prof")
	//fmt.Println(err)
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	GetLog().Printer().SetOutput(new(bytes.Buffer))
	log.SetOutput(new(bytes.Buffer))
	for i:=0; i < b.N; i++ {
		Debugf("i=%d", i)
		//log.Printf("i=%d", i)
	}
}