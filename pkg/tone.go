package synth

import (
	"math"
)

func GenerateTone(freq, dur float64, frameRate int, amp float64) []int {
	if freq == 0 {
		return make([]int, int(float64(frameRate)*dur))
	}
	// TODO : Put incomplete cycles back in, accept param for specifiying initial phase
	frameCount := int(float64(frameRate) * dur)
	oneCycleFrameCount := int(float64(frameRate) / freq)
	tone := make([]int, 0)

	// Calculate one single cycle
	oneCycle := []int{}
	for n := 0; n < oneCycleFrameCount; n++ {
		sample := math.Sin(float64(n) * float64(2*math.Pi) / float64(oneCycleFrameCount))
		sample *= math.Pow(2, 15) * amp
		oneCycle = append(oneCycle, int(sample))
	}

	// Count number of complete cycles, and the length of incomplete cycle
	// TODO : Truncation of frame count means we may be missing frames, fix this
	cycleCount, incompleteCycle := frameCount/oneCycleFrameCount, frameCount%oneCycleFrameCount
	for n := 0; n < cycleCount; n++ {
		tone = append(tone, oneCycle...)
	}

	// Append incomplete cycle
	tone = append(tone, oneCycle[:incompleteCycle]...)

	return tone
}
