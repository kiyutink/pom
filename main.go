package main

import (
	"time"

	"github.com/getlantern/systray"
)

const (
	configWorkDuration = time.Minute * 25
	configRestDuration = time.Minute * 5
	r                  = configWorkDuration
	m                  = work
)

func main() {
	s := &sound{
		filePath: "assets/sms-tone.mp3",
	}
	s.init()
	p := pom{
		remaining: r,
		mode:      m,
		sound:     s,
	}
	systray.Run(p.init, func() {})
}
