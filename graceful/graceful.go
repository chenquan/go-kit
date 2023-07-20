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

package graceful

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chenquan/go-kit/log"
)

const waitTime = 5500 * time.Millisecond

func init() {
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM)

		for {
			v := <-signals
			switch v {
			case syscall.SIGTERM:
				gracefulStop(signals)
			default:
				_ = log.Log(log.LevelError, "Unknown signal", log.FieldKv("signal", v.String()))
			}
		}
	}()
}

func gracefulStop(signals chan os.Signal) {
	signal.Stop(signals)

	_ = log.Log(log.LevelInfo, "Got signal SIGTERM, shutting down...")

	go func() {
		Shutdown()
		_ = log.Log(log.LevelInfo, "Shutdown complete")
	}()

	time.Sleep(waitTime)
	_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
}
