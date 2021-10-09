package automaton

type Automaton interface {
	Step()
}

type SimpleAutomaton2D interface {
	Automaton
	GetData() [][]bool
}

