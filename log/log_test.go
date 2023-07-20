/*
 *    Copyright 2023 chenquan
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

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
