package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("test the functionality", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if sleeper.Calls != 3 {
			t.Errorf("3 calls expected from the sleep, got %d", sleeper.Calls)
		}
	})

	t.Run("test the behavior", func(t *testing.T) {
		sleeper := &SpyCountdownOperation{}
		Countdown(sleeper, sleeper)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, sleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, sleeper.Calls)
		}

	})
}
