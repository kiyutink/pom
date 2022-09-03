package main

import (
	"github.com/getlantern/systray"
)

type button struct {
	title   string
	tooltip string
	handler func()

	mItem *systray.MenuItem
}

func (b *button) init() {
	b.mItem = systray.AddMenuItem(b.title, b.tooltip)
	go func() {
		for range b.mItem.ClickedCh {
			go func() {
				b.handler()
			}()
		}
	}()
}

func (b *button) show() {
	b.mItem.Show()
}

func (b *button) hide() {
	b.mItem.Hide()
}
