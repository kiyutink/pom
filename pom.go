package main

import (
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
	countdown   *countdown
	startButton *button
	pauseButton *button
	resetButton *button
	quitButton  *button
	sound       *sound
}

func (p *pom) init() {
	systray.SetTitle(p.remaining.Round(time.Second).String())
	p.startButton = &button{
		title:   "Start",
		tooltip: "Start the timer",
		handler: func() {
			p.countdown = &countdown{
				dur: p.remaining,
				per: time.Second,
				handler: func(r time.Duration) {
					p.remaining = r
					systray.SetTitle(p.remaining.Round(time.Second).String())
				},
			}
			p.setActive()
			err := p.countdown.start()
			if err != errStopped {
				if p.mode == work {
					p.remaining = configRestDuration
				} else {
					p.remaining = configWorkDuration
				}
				p.toggleMode()
				p.sound.play()
			}
			p.setPaused()
			systray.SetTitle(p.remaining.Round(time.Second).String())
		},
	}
	p.startButton.init()

	p.pauseButton = &button{
		title:   "Pause",
		tooltip: "Pause the timer",
		handler: func() {
			p.setPaused()
			p.countdown.stop()
		},
	}
	p.pauseButton.init()
	p.pauseButton.hide()

	p.resetButton = &button{
		title:   "Reset",
		tooltip: "Reset the timer",
		handler: func() {
			p.countdown.stop()
			p.setPaused()
			p.remaining = r
			systray.SetTitle(p.remaining.Round(time.Second).String())
		},
	}
	p.resetButton.init()

	p.quitButton = &button{
		title:   "Quit",
		tooltip: "Quit the app",
		handler: func() {
			systray.Quit()
		},
	}
	p.quitButton.init()
}

func (p *pom) toggleMode() {
	p.mode = map[mode]mode{
		work: rest,
		rest: work,
	}[p.mode]
}

func (p *pom) setActive() {
	p.pauseButton.show()
	p.startButton.hide()
}

func (p *pom) setPaused() {
	p.pauseButton.hide()
	p.startButton.show()
}
