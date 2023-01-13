package logging

import (
	"bytes"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewSimpleLogger(DEBUG, SimpleFormatter{
		MsgFormat:  "$l [$L] :: $m",
		TimeFormat: "",
		LineEnd:    "",
	}, buf)
	logger.Info("hallo")

	if buf.String() != "[*] [INFO] :: hallo" {
		t.Errorf("buf: %s", buf.String())
	}
}
