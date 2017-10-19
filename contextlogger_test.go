package logging

import (
	"testing"
	"bytes"
)

func TestDebug(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", 0)
	logger.Debug("name=", "elvis");

	if actual := buff.String(); actual != "DEBUG m=TestDebug name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestDebugf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", 0)
	logger.Debugf("name=%v", "elvis");

	if actual := buff.String(); actual != "DEBUG m=TestDebugf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestInfo(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", 0)
	logger.Info("name=", "elvis");

	if actual := buff.String(); actual != "INFO m=TestInfo name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestInfof(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", 0)
	logger.Infof("name=%v", "elvis");

	if actual := buff.String(); actual != "INFO m=TestInfof name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestWarn(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", 0)
	logger.Warning("name=", "elvis");

	if actual := buff.String(); actual != "WARNING m=TestWarn name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestWarnf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", 0)
	logger.Warningf("name=%v", "elvis");

	if actual := buff.String(); actual != "WARNING m=TestWarnf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestError(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", 0)
	logger.Error("name=", "elvis");

	if actual := buff.String(); actual != "ERROR m=TestError name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}

func TestErrorf(t *testing.T){

	buff := new(bytes.Buffer)
	logger := New(buff, "", 0)
	logger.Errorf("name=%v", "elvis");

	if actual := buff.String(); actual != "ERROR m=TestErrorf name=elvis\n" {
		t.Errorf("log format not expected, full=%s, actual=%s", buff.String(), actual)
	}
}
