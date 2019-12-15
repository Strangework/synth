package main

import (
	"fmt"
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
	pcmData := synth.GenerateTone(400, 5, FRAME_RATE, AMPLITUDE)
	fmt.Println(len(pcmData))
	pcmData = synth.AddReverb(pcmData, 0.1, 0.5, FRAME_RATE)
	fmt.Println(len(pcmData))

	tone := audio.IntBuffer{
		&audio.Format{CHANNEL_COUNT, FRAME_RATE},
		pcmData,
		BIT_DEPTH,
	}

	out, err := os.Create("reverb.wav")
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
