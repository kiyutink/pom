package main

import (
	"time"

	"github.com/getlantern/systray"
)

// https://google.com

const (
	configWorkDuration = time.Second * 5
	configRestDuration = time.Second * 3
	r                  = time.Second * 5
	m                  = work
)

func main() {
	p := pom{
		remaining: r,
		mode:      m,
	}
	systray.Run(p.init, func() {})
}
