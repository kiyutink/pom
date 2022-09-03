package main

import (
	"time"

	"github.com/getlantern/systray"
)

const (
	r = time.Minute * 25
	m = work
)

func main() {
	p := pom{
		remaining: r,
		mode:      m,
	}
	systray.Run(p.init, func() {})
}
