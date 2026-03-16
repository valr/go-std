package syncx

import (
	"sync"
	"time"
)

// WaitTimeout waits for the WaitGroup for the specified amount of time.
// Returns true if waiting timed out, false if waiting completed successfully.
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan struct{})

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	select {
	case <-ch:
		return false
	case <-time.After(timeout):
		return true
	}
}
