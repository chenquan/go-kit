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
