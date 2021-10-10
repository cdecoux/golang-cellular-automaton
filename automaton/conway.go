package automaton

/*
	https://en.wikipedia.org/wiki/Conway's_Game_of_Life
 */

type conwayAutomaton struct {
	SimpleAutomaton2D
	Data [][]bool
}

func NewConwayAutomaton(x, y int) *conwayAutomaton {
	// Initialzie the slice of data to the grid size.
	data := make([][]bool, x)
	for i := range data {
		data[i] = make([]bool, y)
	}

	automaton := &conwayAutomaton{
		Data: data,
	}

	return automaton
}


func (self *conwayAutomaton) Step()  {
	self.iterate()
}


func (self *conwayAutomaton) GetData() [][]bool {
	return self.Data
}

/*
	Helper Functions
 */

func (self *conwayAutomaton) iterate() {

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
			stagedData[i][j] = self.getCellUpdate(i, j)
		}
	}

	self.Data = stagedData
}

// https://en.wikipedia.org/wiki/Conway's_Game_of_Life#Rules
func (self *conwayAutomaton) getCellUpdate(x, y int) bool {
	neighbors := self.getNeighbors(x, y)
	neighborCount := 0

	for _, neighbor := range neighbors {
		if neighbor {
			neighborCount++
		}
	}

	if self.Data[x][y] {
		// If cell is "alive", it will die due to underpopulation or overcrowding with fewer than 2 or greater than 3 neighbors respectively
		if neighborCount < 2 || neighborCount > 3 {return false} else {return true}
	} else {
		// If cell is "dead", exactly three live neighbors turns this one to live
		if neighborCount == 3 {return true} else {return false}
	}


}

func (self *conwayAutomaton) getNeighbors(x, y int) [8]bool {
	/*
		Neighbors cells are directly orthogonally or diagonally adjacent
		The following relative coordinates will be used to retrieve the neighbors
		[x - 1, y - 1]	[x, y - 1]	[x + 1, y - 1]
		[x - 1, y]		[x, y]		[x + 1, y]
		[x - 1, y + 1]	[x, y + 1]	[x + 1, y + 1]

		Neighbors that are outside of the bounds will be treated as walls
	 */

	// Get Bounds
	xMax := len(self.Data)
	yMax := len(self.Data[0])

	getNeighborAt := func(x, y int) bool {
		if x < 0 || y < 0 || x >= xMax || y >= yMax {
			return false
		} else {
			return self.Data[x][y]
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