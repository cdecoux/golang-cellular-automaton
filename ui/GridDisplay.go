package ui

import (
	ui "github.com/gizak/termui/v3"
	"golang-cellular-automaton/automaton"
	"image"
)

type gridDisplay struct {
	ui.Block
	Automaton automaton.SimpleAutomaton2D
	CellRune rune
	CellStyle ui.Style
}

func NewGridDisplay(automaton automaton.SimpleAutomaton2D, width, height int) *gridDisplay {
	self := &gridDisplay{
		Block: *ui.NewBlock(),
		CellRune: 0,
		CellStyle : ui.NewStyle(
			ui.StyleClear.Fg,
			ui.Color(8),
			ui.StyleClear.Modifier,
		),
		Automaton: automaton,
	}
	self.SetRect(0, 0, width, height)



	//// Initialzie the slice of data to the grid size.
	//data := make([][]bool, self.Inner.Dx())
	//for i := range data {
	//	data[i] = make([]bool, self.Inner.Dy())
	//}
	return self
}

func (self *gridDisplay) renderData(buf *ui.Buffer)  {
	data := self.Automaton.GetData()

	for i := range data {
		for j := range data[i] {
			if data[i][j] {
				point := image.Point{X: i, Y: j}.Add(self.Inner.Min)
				if point.In(self.Inner) {
					buf.SetCell(
						ui.NewCell(self.CellRune, self.CellStyle),
						image.Point{X: i, Y: j}.Add(self.Inner.Min),
					)
				}
			}
		}
	}

}

func (self *gridDisplay) Update()  {
	self.Automaton.Step()
}

func (self *gridDisplay) Draw(buf *ui.Buffer) {
	self.Block.Draw(buf)

	self.renderData(buf)
}