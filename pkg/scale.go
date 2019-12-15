package synth

import (
	"math"
)

type LetterNote int

const (
	C  LetterNote = 0
	CS LetterNote = 1
	D  LetterNote = 2
	DS LetterNote = 3
	E  LetterNote = 4
	F  LetterNote = 5
	FS LetterNote = 6
	G  LetterNote = 7
	GS LetterNote = 8
	A  LetterNote = 9
	AS LetterNote = 10
	B  LetterNote = 11
)

type Note struct {
	Letter LetterNote
	Octave int
}

func NewNoteByIndex(index int) Note {
	var n Note
	n.Letter = LetterNote(index % 12)
	n.Octave = index / 12
	return n
}

func NewNote(letter LetterNote, octave int) Note {
	var n Note
	n.Letter = letter
	n.Octave = octave
	return n
}

func (n Note) GetIndex() int {
	return n.Octave*12 + int(n.Letter)
}

type EqualTemperament struct {
	Reference Note
	RefFreq   float64
	Delta     float64
}

func NewEqualTemperament(reference Note, refFreq float64, delta float64) *EqualTemperament {
	et := new(EqualTemperament)
	et.Reference = reference
	et.RefFreq = refFreq
	et.Delta = delta
	return et
}
func (et EqualTemperament) GetFrequency(note Note) float64 {
	freq := et.RefFreq * math.Pow(et.Delta, float64(note.GetIndex()-et.Reference.GetIndex()))
	return freq
}
