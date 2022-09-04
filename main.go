package main

import (
	"time"

	"github.com/getlantern/systray"
)

const (
	configWorkDuration = time.Second * 5
	configRestDuration = time.Second * 3
	r                  = time.Second * 5
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
