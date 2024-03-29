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
	mode             mode
	remaining        time.Duration
	countdown        *countdown
	startButton      *button
	pauseButton      *button
	resetButton      *button
	quitButton       *button
	toggleModeButton *button
	sound            *sound
}

func (p *pom) init() {
	p.render()
	p.startButton = &button{
		title:   "Start",
		tooltip: "Start the timer",
		handler: func() {
			p.countdown = &countdown{
				dur: p.remaining,
				per: time.Second,
				handler: func(r time.Duration) {
					p.remaining = r
					p.render()
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
			p.render()
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
			if p.countdown != nil {
				p.countdown.stop()
			}
			p.setPaused()
			p.remaining = r
			p.render()
		},
	}
	p.resetButton.init()

	p.toggleModeButton = &button{
		title:   "Toggle mode",
		tooltip: "Toggle mode and reset the timer",
		handler: func() {
			if p.countdown != nil {
				p.countdown.stop()
			}
			p.setPaused()
			p.toggleMode()
			if p.mode == work {
				p.remaining = configWorkDuration
			} else {
				p.remaining = configRestDuration
			}
			p.render()
		},
	}
	p.toggleModeButton.init()

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

func (p *pom) render() {
	emoji := "👨‍💻"
	if p.mode == rest {
		emoji = "🌴"
	}
	text := fmt.Sprintf("%v %v", emoji, p.remaining.Round(time.Second).String())
	systray.SetTitle(text)
}
