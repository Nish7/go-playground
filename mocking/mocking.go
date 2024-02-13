package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpyCountdownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}

	fmt.Fprint(w, "Go!")
}

func main() {
	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)
}
