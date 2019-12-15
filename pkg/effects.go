package synth

import (
	"math"
)

func AddReverb(wave []int, rev_time, decay_ratio float64, frameRate int) []int {
	decaying_wave := make([]int, len(wave))
	copy(decaying_wave, wave)

	// Determine how many reverberations are needed before sound is 60dB less
	decay_steps := 1
	for math.Pow(decay_ratio, float64(decay_steps)) > 0.001 {
		decay_steps++
	}

	// Resize wave to contain reverberations
	final_size := int(float64(frameRate*decay_steps)*rev_time) + len(wave)
	for len(wave) < final_size {
		wave = append(wave, 0)
	}

	for step := 0; step < decay_steps; step++ {
		// Calculate reverb wave
		rev_start := int(float64(frameRate*step) * rev_time)
		for n := 0; n < len(decaying_wave); n++ {
			decaying_wave[n] = int(float64(decaying_wave[n]) * decay_ratio)
			wave[rev_start+n] += decaying_wave[n]
		}
	}

	return wave
}
