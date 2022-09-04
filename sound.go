package main

import (
	"io"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type sound struct {
	file     io.ReadCloser
	streamer beep.StreamSeekCloser
}

func (s *sound) init() {
	
	streamer, format, err := mp3.Decode(s.file)
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
