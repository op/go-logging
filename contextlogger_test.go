package logging

import (
	"testing"
	"log"
	"bytes"
)

func TestDebugf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", log.LstdFlags)
	logger.Debugf("name=%v", "elvis");

	actual := buff.String()[20:]
	if actual != "DEBUG m=TestDebugf name=elvis\n" {
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
