package automaton

/*
	https://en.wikipedia.org/wiki/Conway's_Game_of_Life
 */

// https://en.wikipedia.org/wiki/Conway's_Game_of_Life#Rules
func ConwayCellUpdateRule(self simpleAutomaton2D, x, y int) bool {
	neighbors := self.getNeighbors(x, y)
	neighborCount := 0

	for _, neighbor := range neighbors {
		if neighbor {
			neighborCount++
		}
	}

	if (self.Data)[x][y] {
		// If cell is "alive", it will die due to underpopulation or overcrowding with fewer than 2 or greater than 3 neighbors respectively
		if neighborCount < 2 || neighborCount > 3 {return false} else {return true}
	} else {
		// If cell is "dead", exactly three live neighbors turns this one to live
		if neighborCount == 3 {return true} else {return false}
	}

}