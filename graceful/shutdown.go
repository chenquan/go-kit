package graceful

import (
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/chenquan/go-kit/log"
)

var (
	shutdownListeners = new(listenerManager)
)

type listenerManager struct {
	lock      sync.Mutex
	waitGroup sync.WaitGroup
	listeners []func()
}

// AddShutdownListener adds fn as a shutdown listener.
// The returned func can be used to wait for fn getting called.
func AddShutdownListener(fn func()) (waitForCalled func()) {
	return shutdownListeners.addListener(fn)
}

func (lm *listenerManager) addListener(fn func()) (waitForCalled func()) {
	lm.waitGroup.Add(1)

	lm.lock.Lock()
	lm.listeners = append(lm.listeners, func() {
		defer lm.waitGroup.Done()
		fn()
	})
	lm.lock.Unlock()

	return func() {
		lm.waitGroup.Wait()
	}
}

func (lm *listenerManager) notifyListeners() {
	lm.lock.Lock()
	defer lm.lock.Unlock()

	wg := sync.WaitGroup{}
	for _, listener := range lm.listeners {
		listener := listener
		wg.Add(1)
		go func() {
			defer func() {
				if e := recover(); e != nil {
					_ = log.Log(log.LevelError, fmt.Sprintf("%v\n%s", e, string(debug.Stack())))
				}

				wg.Done()
			}()
			listener()
		}()
	}

	wg.Wait()
}

func Shutdown() {
	shutdownListeners.notifyListeners()
}
