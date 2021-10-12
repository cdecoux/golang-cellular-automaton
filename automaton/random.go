package automaton

import (
	"math/rand"
	"time"
)

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