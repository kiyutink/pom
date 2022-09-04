package main

import (
	"errors"
	"time"
)

type countdown struct {
	dur     time.Duration
	per     time.Duration
	handler func(t time.Duration)

	ticker  *time.Ticker
	started time.Time
	done    chan struct{}
}

var errStopped = errors.New("countdown stopped")

// start will block until time runs out or until stop is called. TODO: doesn't return when time runs out
// TODO: this should accept context
func (c *countdown) start() error {
	c.ticker = time.NewTicker(c.per)
	defer c.ticker.Stop()
	c.started = time.Now()
	c.done = make(chan struct{})
	for {
		select {
		case t := <-c.ticker.C:
			passed := t.Sub(c.started)
			c.handler(c.dur - passed)
			if passed >= c.dur {
				return nil
			}

		case <-c.done:
			return errStopped
		}
	}
}

// this doesn't make the channel close
func (c *countdown) stop() {
	c.ticker.Stop()
	c.done <- struct{}{}
}
