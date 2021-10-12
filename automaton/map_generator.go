package automaton

/*
	https://en.wikipedia.org/wiki/Conway's_Game_of_Life
*/

type mapGenAutomaton struct {
	SimpleAutomaton2D
	Data [][]bool
}

func NewMapGenAutomaton(x, y int) *mapGenAutomaton {
	// Initialzie the slice of data to the grid size.
	data := make([][]bool, x)
	for i := range data {
		data[i] = make([]bool, y)
	}

	automaton := &mapGenAutomaton{
		Data: data,
	}

	return automaton
}


func (self *mapGenAutomaton) Step()  {
	UpdateCells(self)
}


func (self *mapGenAutomaton) GetData() [][]bool {
	return self.Data
}

func (self *mapGenAutomaton) SetData(data [][]bool)  {
	self.Data = data
}

// https://en.wikipedia.org/wiki/Conway's_Game_of_Life#Rules
func (self *mapGenAutomaton) getCellUpdate(x, y int) bool {
	neighbors := getNeighbors(self, x, y)
	neighborCount := 0

	for _, neighbor := range neighbors {
		if neighbor {
			neighborCount++
		}
	}

	if neighborCount > 3 {
		return false
	} else {return true}

}