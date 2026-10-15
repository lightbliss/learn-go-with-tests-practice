package main

import (
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(d time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func main() {
	ds := new(DefaultSleeper)
	_, err := Countdown(os.Stdout, ds)
	if err != nil {
		log.Fatal(err)
	}

	cs := &ConfigurableSleeper{1 * time.Second, time.Sleep}

	_, err = Countdown(os.Stdout, cs)
	if err != nil {
		log.Fatal(err)
	}
}

func Countdown(out io.Writer, s Sleeper) (int, error) {
	var total int
	for i := range countdownFrom(countdownStart) {
		n, err := fmt.Fprintln(out, i)
		if err != nil {
			log.Println(err)
		}
		total += n
		s.Sleep()
	}
	n, err := fmt.Fprintln(out, finalWord)
	if err != nil {
		log.Println(err)
	}
	total += n

	return total, nil
}

func countdownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
