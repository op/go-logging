package logging

import (
	"testing"
	"log"
	"bytes"
)

func TestDebug(t *testing.T){

	minSize := 20

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Debug("name=", "elvis");

	if buff.Len() < minSize {
		t.Errorf("log must have %d at least, actual=%s", minSize, buff.String())
	}
	if actual := buff.String()[minSize:]; actual != "DEBUG m=TestDebug name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestDebugf(t *testing.T){

	minSize := 20

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Debugf("name=%v", "elvis");

	if buff.Len() < minSize {
		t.Errorf("log must have %d at least, actual=%s", minSize, buff.String())
	}
	if actual := buff.String()[minSize:]; actual != "DEBUG m=TestDebugf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestInfo(t *testing.T){

	minSize := 20

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Info("name=", "elvis");

	if buff.Len() < minSize {
		t.Errorf("log must have %d at least, actual=%s", minSize, buff.String())
	}
	if actual := buff.String()[minSize:]; actual != "INFO m=TestInfo name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestInfof(t *testing.T){

	minSize := 20

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Infof("name=%v", "elvis");

	if buff.Len() < minSize {
		t.Errorf("log must have %d at least, actual=%s", minSize, buff.String())
	}
	if actual := buff.String()[minSize:]; actual != "INFO m=TestInfof name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestWarn(t *testing.T){

	minSize := 20

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Warning("name=", "elvis");

	if buff.Len() < minSize {
		t.Errorf("log must have %d at least, actual=%s", minSize, buff.String())
	}
	if actual := buff.String()[minSize:]; actual != "WARNING m=TestWarn name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestWarnf(t *testing.T){

	minSize := 20

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Warningf("name=%v", "elvis");

	if buff.Len() < minSize {
		t.Errorf("log must have %d at least, actual=%s", minSize, buff.String())
	}
	if actual := buff.String()[minSize:]; actual != "WARNING m=TestWarnf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestError(t *testing.T){

	minSize := 20

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Error("name=", "elvis");

	if buff.Len() < minSize {
		t.Errorf("log must have %d at least, actual=%s", minSize, buff.String())
	}
	if actual := buff.String()[minSize:]; actual != "ERROR m=TestError name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestErrorf(t *testing.T){

	minSize := 20

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Errorf("name=%v", "elvis");

	if buff.Len() < minSize {
		t.Errorf("log must have %d at least, actual=%s", minSize, buff.String())
	}
	if actual := buff.String()[minSize:]; actual != "ERROR m=TestErrorf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}


//func TestNewLogWithLogger(t *testing.T) {
//
//	buff := bytes.Buffer{}
//
//	logger := MustGetLogger("main")
//	SetFormatter(MustStringFormatter(`%{level:.3s} %{message}`))
//	SetBackend(NewLogBackend(&buff, "", 0))
//
//	NewLogWithLogger(logger).Info("x1")
//	NewLogWithLogger(logger).Infof("x2")
//
//	expectedLog := "INF id=1, m=TestNewLogWithLogger x1\nINF id=2, m=TestNewLogWithLogger x2\n"
//	if buff.String() != expectedLog {
//		t.Fail()
//		t.Error("current: " + buff.String())
//		t.Error("expected: " + expectedLog)
//	}
//	fmt.Println(buff.String())
//
//}
//
//func TestNewLog(t *testing.T) {
//
//	buff := bytes.Buffer{}
//
//	SetFormatter(MustStringFormatter(`%{level:.3s} %{message}`))
//	SetBackend(NewLogBackend(&buff, "", 0))
//
//	NewLog(NewContext()).Info("y1")
//	NewLog(NewContext()).Infof("y2")
//
//	expectedLog := "INF id=3, m=TestNewLog y1\nINF id=4, m=TestNewLog y2\n"
//
//
//	if buff.String() != expectedLog {
//		t.Fail()
//		t.Error("current: " + buff.String())
//		t.Error("expected: " + expectedLog)
//	}
//	fmt.Println(buff.String())
//}
//
//
