package automaton

import (
	"math/rand"
	"time"
)

/*
	Randomizer Automaton
 */

type randomAutomaton struct {
	SimpleAutomaton2D
	Data [][]bool
}

func NewRandomAutomaton(x, y int) *randomAutomaton {
	// Initialzie the slice of data to the grid size.
	data := make([][]bool, x)
	for i := range data {
		data[i] = make([]bool, y)
	}

	automaton := &randomAutomaton{
		Data: data,
	}

	automaton.FillRandom()

	return automaton
}

func (self *randomAutomaton) Step()  {
	boolgen := NewBoolGenerator()

	for i := range self.Data {
		for j := range self.Data[i] {
			self.Data[i][j] = boolgen.Bool()
		}
	}
}


func (self *randomAutomaton) GetData() [][]bool {
	return self.Data
}

/*
	Randomizer Helper Functions
 */

func (self *randomAutomaton) FillRandom()  {
	boolgen := NewBoolGenerator()

	data := self.Data

	for i := range data {
		for j := range data[i] {
			data[i][j] = boolgen.Bool()
		}
	}
}

/*
	https://stackoverflow.com/a/45031417/14915694
 */

type boolgenerator struct {
	src       rand.Source
	cache     int64
	remaining int
}

func NewBoolGenerator() *boolgenerator {
	return &boolgenerator{src: rand.NewSource(time.Now().UnixNano())}
}

func (b *boolgenerator) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}