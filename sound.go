package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

type sound struct {
	player *audio.Player
}

func newSound(filepath string, loop bool) *sound {
	var streamType io.ReadSeeker
	audioBytes, err := audioFS.ReadFile(filepath)
	if err != nil {
		panic("can't open embedded audio")
	}
	buffer := bytes.NewReader(audioBytes)
	br, err := vorbis.Decode(audioContext, buffer)
	if err != nil {
		panic(fmt.Sprintf("couldn't decode audio:%v", err))
	}
	streamType = br
	if loop {
		streamType = audio.NewInfiniteLoop(br, br.Length())
	}
	audioPlayer, err := audio.NewPlayer(audioContext, streamType)
	if err != nil {
		panic("couldn't create player")
	}
	return &sound{
		player: audioPlayer,
	}

}

func (s *sound) play() {
	s.player.Play()
}

func (s *sound) stop() {
	s.player.Pause()
	s.player.Rewind()
}
