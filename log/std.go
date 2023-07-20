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
	"fmt"
	stdlog "log"
	"strings"
)

func init() {
	SetLogger(newStd())
}

type std struct {
	log *stdlog.Logger
}

func newStd() Logger {
	return &std{log: stdlog.Default()}
}

func (l std) Log(level Level, s string, fields ...Field) error {
	values := make([]string, 0, 1+len(s)+len(fields))
	values = append(values, level.String())

	if len(s) != 0 {
		values = append(values, s)
	}

	if len(fields) != 0 {
		values = append(values, buildField(fields...)...)
	}

	l.log.Println(strings.Join(values, "\t"))
	return nil
}

func buildField(fields ...Field) []string {
	values := make([]string, 0, len(fields))
	for _, f := range fields {
		values = append(values, fmt.Sprintf("%s=%+v", f.Key(), f.Value()))
	}

	return values
}
