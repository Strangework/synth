package main

import (
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"math"
	"os"
	"strangework/synth/pkg"
)

var AMPLITUDE = 0.8
var FRAME_RATE = 48000
var BIT_DEPTH = 16
var CHANNEL_COUNT = 1

func main() {
	var refFreq float64 = 440
	refDelta := math.Pow(2, 1.0/12)
	fmt.Printf("Reference frequency: %f\nDelta: %e\n",
		refFreq, refDelta)

	// A4, assuming index 0 refers to C0, instead of the typical A0
	//aFour := synth.NewNote(synth.A, 4)
	refNote := synth.NewNoteByIndex(57) // Identical to A4
	et := synth.NewEqualTemperament(refNote, refFreq, refDelta)

	fmt.Println("Sweeping 12-TET starting from C0")
	pcmData := make([]int, 0)
	for n := 0; n < 120; n++ {
		note := synth.NewNoteByIndex(n)
		freq := et.GetFrequency(note)
		fmt.Printf("C0 + %d: %f\n", n, freq)
		tone := synth.GenerateTone(freq, 0.5, FRAME_RATE, AMPLITUDE)
		pcmData = append(pcmData, tone...)
	}

	tone := audio.IntBuffer{
		&audio.Format{CHANNEL_COUNT, FRAME_RATE},
		pcmData,
		BIT_DEPTH,
	}

	out, err := os.Create("12-tet.wav")
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
