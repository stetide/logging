package logging_test

import (
	"bytes"
	"cloud/pkg/logging"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := logging.NewSimpleLogger(logging.DEBUG, logging.SimpleFormatter{
		"$l [$L] :: $m",
		"",
		"",
	}, buf)
	logger.Info("hallo")

	if buf.String() != "[*] [INFO] :: hallo" {
		t.Errorf("buf: %s", buf.String())
	}
}
