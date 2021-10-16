package ui

import (
	ui "github.com/gizak/termui/v3"
	"golang-cellular-automaton/automaton"
	"image"
	"math/rand"
)

type complexGridDisplay struct {
	ui.Block
	Automaton automaton.ComplexAutomaton2D
	CellRune rune
	CellStyle ui.Style
	DefaultStyle ui.Style
	Colors []ui.Color
}

func NewComplexGridDisplay(automaton automaton.ComplexAutomaton2D, width, height int) *complexGridDisplay {
	self := &complexGridDisplay{
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

	for i := 0; i < len(automaton.GetTypes()) - 1; i++ {
		self.Colors = append(self.Colors, ui.Color(rand.Intn(255)))
	}

	self.SetRect(0, 0, width, height)

	return self
}

func (self *complexGridDisplay) renderData(buf *ui.Buffer)  {
	data := self.Automaton.GetData()

	for i := range data {
		for j := range data[i] {
			point := image.Point{X: i, Y: j}.Add(self.Inner.Min)
			cellRune := ' '
			cellStyle := self.DefaultStyle

			if data[i][j] > 0{
				cellRune = self.CellRune
				cellStyle = self.CellStyle
				cellStyle.Fg = ui.SelectColor(self.Colors, data[i][j])
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

func (self *complexGridDisplay) Update()  {
	self.Automaton.Step()
}

func (self *complexGridDisplay) UpdateMouseEvent(mouse ui.Mouse) {
	x := mouse.X
	y := mouse.Y

	point := image.Point{
		X: x - self.Inner.Min.X,
		Y: y - self.Inner.Min.Y,
	}

	if point.In(self.Inner) {
		data := self.Automaton.GetData()
		types := self.Automaton.GetTypes()
		data[point.X][point.Y] = (data[point.X][point.Y] + 1) % len(types)
	}
}

func (self *complexGridDisplay) Draw(buf *ui.Buffer) {
	self.Block.Draw(buf)

	self.renderData(buf)
}