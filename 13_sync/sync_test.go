package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, counter.Value())
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantCount)
	})
}

func assertCounter(t *testing.T, c *Counter, value int) {
	t.Helper()

	if c.Value() != value {
		t.Errorf("got %d, want %d", c.Value(), value)
	}
}
