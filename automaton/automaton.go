package automaton

import (
	"math/rand"
	"time"
)

type Automaton interface {
	Step()
}

type SimpleAutomaton2D interface {
	Automaton
	GetData() [][]bool
	SetData([][]bool)
	getCellUpdate(x, y int) bool
}


/*
	Helper Functions
 */

func FillRandom(automaton SimpleAutomaton2D)  {
	boolgen := NewBoolGenerator()

	data := automaton.GetData()

	for i := range data {
		for j := range data[i] {
			data[i][j] = boolgen.Bool()
		}
	}
}

func getNeighbors(automaton SimpleAutomaton2D, x, y int) [8]bool {
	/*
		Neighbors cells are directly orthogonally or diagonally adjacent
		The following relative coordinates will be used to retrieve the neighbors
		[x - 1, y - 1]	[x, y - 1]	[x + 1, y - 1]
		[x - 1, y]		[x, y]		[x + 1, y]
		[x - 1, y + 1]	[x, y + 1]	[x + 1, y + 1]

		Neighbors that are outside of the bounds will be treated as walls
	*/
	data := automaton.GetData()

	// Get Bounds
	xMax := len(data)
	yMax := len(data[0])

	getNeighborAt := func(x, y int) bool {
		if x < 0 || y < 0 || x >= xMax || y >= yMax {
			return false
		} else {
			return data[x][y]
		}
	}

	return [8]bool {
		getNeighborAt(x - 1, y - 1),
		getNeighborAt(x,     y - 1),
		getNeighborAt(x + 1, y - 1),
		getNeighborAt(x - 1, y),
		getNeighborAt(x + 1, y),
		getNeighborAt(x - 1, y + 1),
		getNeighborAt(x,     y + 1),
		getNeighborAt(x + 1, y + 1),
	}
}


func UpdateCells(automaton SimpleAutomaton2D) {

	data := automaton.GetData()

	// Initialzie a new Data block to put our updates in
	x := len(data)
	y := len(data[0])

	stagedData := make([][]bool, x)
	for i := range stagedData {
		stagedData[i] = make([]bool, y)
	}

	// Go through each cell and update in the staging area
	for i := range data {
		for j := range data[i] {
			stagedData[i][j] = automaton.getCellUpdate(i, j)
		}
	}

	automaton.SetData(stagedData)
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