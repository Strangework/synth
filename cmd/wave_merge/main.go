package main

import (
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"os"
	"strangework/synth/pkg"
)

var AMPLITUDE = 0.5
var FRAME_RATE = 48000
var BIT_DEPTH = 16
var CHANNEL_COUNT = 1

func main() {
	toneA := synth.GenerateTone(440, 5, FRAME_RATE, AMPLITUDE)
	toneB := synth.GenerateTone(400, 5, FRAME_RATE, AMPLITUDE)

	for n := 0; n < len(toneA); n++ {
		toneA[n] += toneB[n]
	}
	var pcmData []int = toneA

	tone := audio.IntBuffer{
		&audio.Format{CHANNEL_COUNT, FRAME_RATE},
		pcmData,
		BIT_DEPTH,
	}

	out, err := os.Create("wave-concat.wav")
	if err != nil {
		panic(err)
	}

	e := wav.NewEncoder(out,
		FRAME_RATE,
		BIT_DEPTH,
		CHANNEL_COUNT,
		1) // PCM values

	if err = e.Write(&tone); err != nil {
		panic(err)
	}

	if err = e.Close(); err != nil {
		panic(err)
	}
	out.Close()
}
