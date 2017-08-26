package logging

import (
	"testing"
	"bytes"
	"fmt"
)

func TestNewLog(t *testing.T) {

	buff := bytes.Buffer{}

	logger := MustGetLogger("main")
	SetFormatter(MustStringFormatter(`%{level:.3s} %{message}`))
	SetBackend(NewLogBackend(&buff, "", 0))

	NewLog(logger).Info("x1")
	NewLog(logger).Infof("x2")

	expectedLog := "INF id=1, m=TestNewLog x1\nINF id=2, m=TestNewLog x2\n"
	if buff.String() != expectedLog {
		t.Fail()
		t.Error("current: " + buff.String())
		t.Error("expcted: " + expectedLog)
	}
	fmt.Println(buff.String())

}



