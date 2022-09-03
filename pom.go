package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
)

type mode uint8

const (
	_ mode = iota
	work
	rest
)

type pom struct {
	mode        mode
	remaining   time.Duration
	startButton *button
	pauseButton *button
	resetButton *button
	quitButton  *button
}

func (p *pom) init() {
	systray.SetTitle(p.remaining.Round(time.Second).String())
	p.startButton = &button{
		title:   "Start",
		tooltip: "Start the timer",
		handler: func() {
			fmt.Println("start clicked")
		},
	}
	p.startButton.init()

	p.pauseButton = &button{
		title:   "Pause",
		tooltip: "Pause the timer",
		handler: func() {
			fmt.Println("pause clicked")
		},
	}
	p.pauseButton.init()
	p.pauseButton.hide()

	p.resetButton = &button{
		title:   "Reset",
		tooltip: "Reset the timer",
		handler: func() {
			fmt.Println("reset clicked")
		},
	}
	p.resetButton.init()

	p.quitButton = &button{
		title:   "Quit",
		tooltip: "Quit the app",
		handler: func() {
			fmt.Println("quit clicked")
		},
	}
	p.quitButton.init()
}
