package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times and leaves at it 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3

		assertCounter(t, &counter, want)
	})

	t.Run("run concurrently safely", func(t *testing.T) {
		wantedCounter := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCounter)

		for i := 0; i < wantedCounter; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
