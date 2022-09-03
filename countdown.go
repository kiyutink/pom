package main

import (
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

// start will block until time runs out or until stop is called. TODO: doesn't return when time runs out
// TODO: this should accept context
func (c *countdown) start() {
	c.ticker = time.NewTicker(c.per)
	defer c.ticker.Stop()
	c.started = time.Now()
	c.done = make(chan struct{})
	defer func() { c.done <- struct{}{} }()
	for {
		select {
		case t := <-c.ticker.C:
			passed := t.Sub(c.started)
			c.handler(c.dur - passed)

		case <-c.done:
			return
		}
	}
}

// this doesn't make the channel close
func (c *countdown) stop() {
	c.ticker.Stop()
	c.done <- struct{}{}
}
