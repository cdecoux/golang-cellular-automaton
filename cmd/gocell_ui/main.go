package main

import (
	termui "github.com/gizak/termui/v3"
	log "github.com/sirupsen/logrus"
	"golang-cellular-automaton/automaton"
	"golang-cellular-automaton/ui"
	"golang.org/x/term"
	"time"
)

func main() {
	if err := termui.Init(); err != nil {
		log.Fatal("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	width, height, _ := term.GetSize(0)
	log.Debug("Initializing Terminal Window with Width: %s and Height: %s", width, height)

	gridAutomaton := automaton.NewSimpleAutomaton2D(width, height, automaton.ConwayCellUpdateRule)
	gridAutomaton.FillRandom()


	grid := ui.NewGridDisplay(gridAutomaton, width, height)
	grid.Title = "Golang Cellular Automaton"
	grid.CellRune = '●'
	grid.CellStyle.Fg = termui.Color(214)
	grid.CellStyle.Bg = termui.Color(242)
	grid.DefaultStyle.Bg = termui.Color(239)

	termui.Render(grid)

	tickerCount := 1
	tickerCount++
	uiEvents := termui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
				case "q", "<C-c>":
					return
				case "<Space>", "<Enter>":
					grid.Update()
					termui.Render(grid)
				case "<MouseLeft>":
					grid.UpdateMouseEvent(e.Payload.(termui.Mouse))
					termui.Render(grid)
			}
		case <-ticker:
			//grid.Update()
			termui.Render(grid)
			tickerCount++
		}
	}
}

