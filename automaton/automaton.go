package automaton

type SimpleAutomaton2D interface {
	Step()
	GetData() [][]bool
}

type simpleAutomaton2D struct {
	SimpleAutomaton2D
	Data [][]bool
	cellUpdateRule func(self *simpleAutomaton2D, x, y int) bool
}

func NewSimpleAutomaton2D(width, height int, cellUpdateRule func(self *simpleAutomaton2D, x, y int) bool) *simpleAutomaton2D {
	// Initialzie the slice of data to the grid size.
	data := make([][]bool, width)
	for i := range data {
		data[i] = make([]bool, height)
	}

	return &simpleAutomaton2D{
		Data: data,
		cellUpdateRule: cellUpdateRule,
	}
}



func (self *simpleAutomaton2D) GetData() [][]bool {
	return self.Data
}

func (self *simpleAutomaton2D) Step()  {
	self.UpdateCells()
}

/*
	Helper Functions
 */

func (self *simpleAutomaton2D) FillRandom()  {
	boolgen := NewBoolGenerator()

	data := self.Data

	for i := range data {
		for j := range data[i] {
			data[i][j] = boolgen.Bool()
		}
	}
}

func (self *simpleAutomaton2D)  getNeighbors(x, y int) [8]bool {
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


func (self *simpleAutomaton2D) UpdateCells() {

	// Initialzie a new Data block to put our updates in
	x := len(self.Data)
	y := len(self.Data[0])

	stagedData := make([][]bool, x)
	for i := range stagedData {
		stagedData[i] = make([]bool, y)
	}

	// Go through each cell and update in the staging area
	for i := range self.Data {
		for j := range self.Data[i] {
			stagedData[i][j] = self.cellUpdateRule(self, i, j)
		}
	}

	self.Data = stagedData
}