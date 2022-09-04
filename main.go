package main

import (
	"embed"
	"log"
	"time"

	"github.com/getlantern/systray"
)

const (
	configWorkDuration = time.Minute * 25
	configRestDuration = time.Minute * 5
	r                  = configWorkDuration
	m                  = work
)

//go:embed assets/sms-tone.mp3
var f embed.FS

func main() {
	file, err := f.Open("assets/sms-tone.mp3")
	if err != nil {
		log.Fatal(err)
	}
	s := &sound{
		file: file,
	}
	s.init()
	p := pom{
		remaining: r,
		mode:      m,
		sound:     s,
	}
	systray.Run(p.init, func() {})
}
