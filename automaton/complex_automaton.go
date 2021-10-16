package automaton

import "math/rand"

type ComplexAutomaton2D interface {
	Step()
	GetData() [][]int
	GetTypes() map[string]int
}

type complexAutomaton2D struct {
	ComplexAutomaton2D
	Data [][]int
	Types map[string]int
	cellUpdateRule func(self ComplexAutomaton2D, x, y int) int
}

func NewComplexAutomaton2D(width, height int, names []string, cellUpdateRule func(self ComplexAutomaton2D, x, y int) int) *complexAutomaton2D {
	// Initialzie the slice of data to the grid size.
	data := make([][]int, width)
	for i := range data {
		data[i] = make([]int, height)
	}

	types := make(map[string]int)
	types["Empty"] = 0

	for i, name := range names {
		types[name] = i + 1
	}

	return &complexAutomaton2D{
		Data: data,
		Types: types,
		cellUpdateRule: cellUpdateRule,
	}
}



func (self *complexAutomaton2D) GetData() [][]int {
	return self.Data
}

func (self *complexAutomaton2D) GetTypes() map[string]int {
	return self.Types
}

func (self *complexAutomaton2D) Step()  {
	self.UpdateCells()
}

/*
	Helper Functions
*/

func (self *complexAutomaton2D) FillRandom()  {

	for i := range self.Data {
		for j := range self.Data[i] {
			self.Data[i][j] = rand.Intn(len(self.Types) + 1)
		}
	}
}

func (self *complexAutomaton2D)  getNeighbors(x, y int) [8]int {
	/*
		Neighbors cells are directly orthogonally or diagonally adjacent
		The following relative coordinates will be used to retrieve the neighbors
		[x - 1, y - 1]	[x, y - 1]	[x + 1, y - 1]
		[x - 1, y]		[x, y]		[x + 1, y]
		[x - 1, y + 1]	[x, y + 1]	[x + 1, y + 1]

		Neighbors that are outside of the bounds will be treated as walls
	*/
	data := self.Data

	// Get Bounds
	xMax := len(data)
	yMax := len(data[0])

	getNeighborAt := func(x, y int) int {
		if x < 0 || y < 0 || x >= xMax || y >= yMax {
			return 0
		} else {
			return data[x][y]
		}
	}

	return [8]int {
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


func (self *complexAutomaton2D) UpdateCells() {

	// Initialzie a new Data block to put our updates in
	x := len(self.Data)
	y := len(self.Data[0])

	stagedData := make([][]int, x)
	for i := range stagedData {
		stagedData[i] = make([]int, y)
	}

	// Go through each cell and update in the staging area
	for i := range self.Data {
		for j := range self.Data[i] {
			stagedData[i][j] = self.cellUpdateRule(self, i, j)
		}
	}

	self.Data = stagedData
}