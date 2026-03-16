package syncx

import (
	"sync"
	"testing"
	"time"
)

func TestWaitTimeout(t *testing.T) {
	tests := []struct {
		name     string
		timeout  time.Duration
		waitTime time.Duration
		expected bool
	}{
		{"don't timeout", 100 * time.Millisecond, 10 * time.Millisecond, false},
		{"timeout", 10 * time.Millisecond, 100 * time.Millisecond, true},
		{"zero timeout", 0, 100 * time.Millisecond, true},
		{"negative timeout", -10 * time.Millisecond, 100 * time.Millisecond, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				time.Sleep(tt.waitTime)
				wg.Done()
			}()

			result := WaitTimeout(&wg, tt.timeout)
			if result != tt.expected {
				t.Errorf("WaitTimeout(%v) returned %v; want %v", tt.timeout, result, tt.expected)
			}
		})
	}
}
