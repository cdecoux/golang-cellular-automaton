package main

import (
	termui "github.com/gizak/termui/v3"
	log "github.com/sirupsen/logrus"
	"golang.org/x/term"
	"golang-cellular-automaton/ui"
)

func main() {
	if err := termui.Init(); err != nil {
		log.Fatal("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	width, height, _ := term.GetSize(0)
	log.Debug("Initializing Terminal Window with Width: %s and Height: %s", width, height)


	grid := ui.NewGridDisplay()
	grid.Title = "Test"
	grid.SetRect(0, 0, width, height)

	termui.Render(grid)

	for e := range termui.PollEvents() {
		if e.Type == termui.KeyboardEvent {
			break
		}
	}
}

