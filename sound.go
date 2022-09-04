package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	_ "github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type sound struct {
	filePath string
	streamer beep.StreamSeekCloser
}

func (s *sound) init() {
	f, err := os.Open(s.filePath)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	s.streamer = streamer
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
}

func (s *sound) play() {
	s.streamer.Seek(0)
	speaker.Play(s.streamer)
}
