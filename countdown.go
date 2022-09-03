package main

import "time"

type countdown struct {
	dur    time.Duration
	per    time.Duration
	ticker time.Ticker
}

func newCountdown(dur time.Duration, per time.Duration) *countdown {
	return &countdown{
		dur: dur,
		per: per,
	}
}

// start will block until
func (c *countdown) start() {
}

func (c *countdown) stop() {
}
