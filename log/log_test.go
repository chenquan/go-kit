package log

import (
	"bytes"
	stdlog "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	logger := getLogger()
	buf := bytes.Buffer{}
	m := &std{log: stdlog.New(&buf, "", stdlog.LstdFlags)}
	SetLogger(m)
	defer SetLogger(logger)

	_ = Log(LevelInfo, "info", FieldKv("a", "any"))
	assert.Contains(t, buf.String(), "info")
	assert.Contains(t, buf.String(), "any")
	buf.Reset()

	_ = Log(LevelDebug, "debug")
	assert.Contains(t, buf.String(), "debug")
	buf.Reset()

	_ = Log(LevelError, "error")
	assert.Contains(t, buf.String(), "error")
	buf.Reset()

	_ = Log(LevelWarn, "warn")
	assert.Contains(t, buf.String(), "warn")
	buf.Reset()

	_ = Log(LevelFatal, "fatal")
	assert.Contains(t, buf.String(), "fatal")
	buf.Reset()
}
