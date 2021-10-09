package ui

import (
	ui "github.com/gizak/termui/v3"
	"image"
)

type GridDisplay struct {
	ui.Block
	CellRune rune
	Data [][]bool
}

func NewGridDisplay() *GridDisplay {
	return &GridDisplay{
		Block: *ui.NewBlock(),
		CellRune: ' ',
	}
}

func (self *GridDisplay) renderData(buf *ui.Buffer)  {
	//drawArea := self.Inner // Extrapolate incase this needs to be changed
	cellStyle := ui.NewStyle(
		ui.StyleClear.Fg,
		ui.ColorYellow,
		ui.StyleClear.Modifier,
	)
	point := image.Point{X: 1, Y: 1}
	buf.SetCell(ui.NewCell(self.CellRune, cellStyle), point)

	point = image.Point{X: 1, Y: 2}
	buf.SetCell(ui.NewCell(self.CellRune, cellStyle), point)

	point = image.Point{X: 2, Y: 1}
	buf.SetCell(ui.NewCell(self.CellRune, cellStyle), point)

}

func (self *GridDisplay) Draw(buf *ui.Buffer) {
	self.Block.Draw(buf)

	self.renderData(buf)
}