package main

import (
	termui "github.com/gizak/termui/v3"
	log "github.com/sirupsen/logrus"
	"golang-cellular-automaton/automaton"
	"golang-cellular-automaton/ui"
	"golang.org/x/term"
)

func main() {
	if err := termui.Init(); err != nil {
		log.Fatal("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	width, height, _ := term.GetSize(0)
	log.Debug("Initializing Terminal Window with Width: %s and Height: %s", width, height)

	randAutomaton := automaton.NewRandomAutomaton(width, height)

	grid := ui.NewGridDisplay(randAutomaton, width, height)
	grid.Title = "Golang Cellular Automaton"

	termui.Render(grid)

	for e := range termui.PollEvents() {
		if e.Type == termui.KeyboardEvent {
			break
		}
	}

}

