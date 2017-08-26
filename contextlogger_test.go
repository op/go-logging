package logging

import (
	"testing"
	"bytes"
	"fmt"
)

func TestNewLogWithLogger(t *testing.T) {

	buff := bytes.Buffer{}

	logger := MustGetLogger("main")
	SetFormatter(MustStringFormatter(`%{level:.3s} %{message}`))
	SetBackend(NewLogBackend(&buff, "", 0))

	NewLogWithLogger(logger).Info("x1")
	NewLogWithLogger(logger).Infof("x2")

	expectedLog := "INF id=1, m=TestNewLogWithLogger x1\nINF id=2, m=TestNewLogWithLogger x2\n"
	if buff.String() != expectedLog {
		t.Fail()
		t.Error("current: " + buff.String())
		t.Error("expected: " + expectedLog)
	}
	fmt.Println(buff.String())

}

func TestNewLog(t *testing.T) {

	buff := bytes.Buffer{}

	SetFormatter(MustStringFormatter(`%{level:.3s} %{message}`))
	SetBackend(NewLogBackend(&buff, "", 0))

	NewLog(NewContext()).Info("y1")
	NewLog(NewContext()).Infof("y2")

	expectedLog := "INF id=3, m=TestNewLog y1\nINF id=4, m=TestNewLog y2\n"


	if buff.String() != expectedLog {
		t.Fail()
		t.Error("current: " + buff.String())
		t.Error("expected: " + expectedLog)
	}
	fmt.Println(buff.String())
}


