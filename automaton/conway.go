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
	UpdateCells(self)
}


func (self *conwayAutomaton) GetData() [][]bool {
	return self.Data
}

func (self *conwayAutomaton) SetData(data [][]bool)  {
	self.Data = data
}

// https://en.wikipedia.org/wiki/Conway's_Game_of_Life#Rules
func (self *conwayAutomaton) getCellUpdate(x, y int) bool {
	neighbors := getNeighbors(self, x, y)
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