package ui

import (
	ui "github.com/gizak/termui/v3"
	"image"
)

type GridDisplay struct {
	ui.Block
	CellRune rune
	Data [][]bool
	CellStyle ui.Style
}

func NewGridDisplay(x1, y1, x2, y2 int) *GridDisplay {
	self := &GridDisplay{
		Block: *ui.NewBlock(),
		CellRune: 'Ɵ',
		CellStyle : ui.NewStyle(
			ui.ColorBlack,
			ui.ColorWhite,
			ui.StyleClear.Modifier,
		),
	}
	self.SetRect(x1, y1, x2, y2)



	// Initialzie the slice of data to the grid size.
	data := make([][]bool, self.Inner.Dx())
	for i := range data {
		data[i] = make([]bool, self.Inner.Dy())
	}

	self.Data = data
	return self
}

func (self *GridDisplay) renderData(buf *ui.Buffer)  {
	for i := range self.Data {
		for j := range self.Data[i] {
			if self.Data[i][j] {
				buf.SetCell(
					ui.NewCell(self.CellRune, self.CellStyle),
					image.Point{X: i, Y: j}.Add(self.Inner.Min),
				)
			}
		}
	}

}

func (self *GridDisplay) Draw(buf *ui.Buffer) {
	self.Block.Draw(buf)

	self.renderData(buf)
}