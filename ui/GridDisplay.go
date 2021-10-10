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
	DefaultStyle ui.Style
}

func NewGridDisplay(automaton automaton.SimpleAutomaton2D, width, height int) *gridDisplay {
	self := &gridDisplay{
		Block: *ui.NewBlock(),
		CellRune: 0,
		CellStyle : ui.NewStyle(
			ui.StyleClear.Fg,
			ui.ColorWhite,
			ui.StyleClear.Modifier,
		),
		DefaultStyle : ui.StyleClear,
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
			point := image.Point{X: i, Y: j}.Add(self.Inner.Min)
			cellRune := ' '
			cellStyle := self.DefaultStyle

			if data[i][j] {
				cellRune = self.CellRune
				cellStyle = self.CellStyle
			}

			if point.In(self.Inner) {
				buf.SetCell(
					ui.NewCell(cellRune, cellStyle),
					image.Point{X: i, Y: j}.Add(self.Inner.Min),
				)
			}
		}
	}

}

func (self *gridDisplay) Update()  {
	self.Automaton.Step()
}

func (self *gridDisplay) UpdateMouseEvent(mouse ui.Mouse) {
	x := mouse.X
	y := mouse.Y

	point := image.Point{
		X: x - self.Inner.Min.X,
		Y: y - self.Inner.Min.Y,
	}

	if point.In(self.Inner) {
		data := self.Automaton.GetData()
		data[point.X][point.Y] = !data[point.X][point.Y]
	}
}

func (self *gridDisplay) Draw(buf *ui.Buffer) {
	self.Block.Draw(buf)

	self.renderData(buf)
}