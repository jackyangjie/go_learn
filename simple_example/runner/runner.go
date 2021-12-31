package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	task      []func(int)
}

var ErrTimeOut = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *runner {
	return &runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *runner) Add(task ...func(int)) {
	r.task = append(task)
}

func (r *runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func (r *runner) run() error {
	for id, task := range r.task {
		if r.goInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *runner) goInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
