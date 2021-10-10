package automaton

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

	FillRandom(automaton)

	return automaton
}

func (self *randomAutomaton) GetData() [][]bool {
	return self.Data
}

func (self *randomAutomaton) SetData(data [][]bool)  {
	self.Data = data
}


func (self *randomAutomaton) Step()  {
	FillRandom(self)
}

// https://en.wikipedia.org/wiki/Conway's_Game_of_Life#Rules
func (self *randomAutomaton) getCellUpdate(x, y int) bool {
	return NewBoolGenerator().Bool()

}