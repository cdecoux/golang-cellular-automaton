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