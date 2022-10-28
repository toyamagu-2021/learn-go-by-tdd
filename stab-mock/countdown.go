package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finallWord = "Go!"
const countDownStart = 3
const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type Sleeper interface {
	Sleep()
}

type DefalultSleeper struct{}

func (d *DefalultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countDownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}
	s.Sleep()
	fmt.Fprint(out, finallWord)
}

func main() {
	//s := &DefalultSleeper{}
	s := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, s)
}
