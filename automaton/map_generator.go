package automaton

func MapGenCellUpdateRule(self *simpleAutomaton2D, x, y int) bool {
	neighbors := self.getNeighbors(x, y)
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