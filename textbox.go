package gobless

import (
	"image"
)

type TextBox struct {
	width       int
	height      int
	x           int
	y           int
	Text        string
	Wrap        bool
	Style       Style
	BorderColor Color
	title       string
}

func NewTextBox() *TextBox {
	return &TextBox{
		Style:       DefaultStyle,
		BorderColor: NewColor(0, 255, 255),
	}
}

func (textbox *TextBox) SetTitle(title string) {
	textbox.title = title
}

func (textbox *TextBox) SetX(x int) {
	textbox.x = x
}
func (textbox *TextBox) SetY(y int) {
	textbox.y = y
}
func (textbox *TextBox) SetWidth(w int) {
	textbox.width = w
}
func (textbox *TextBox) SetHeight(h int) {
	textbox.height = h
}

func (textbox *TextBox) GetTiles(gui *GUI) []Tile {

	cells := map[image.Point]Cell{}

	x, y := 0, 0

	textWidth := textbox.width - 2 // remove border dimensions
	textHeight := textbox.height - 2

	for _, r := range []rune(textbox.Text) {
		if !textbox.Wrap && x >= textWidth {
			continue
		}
		switch r {
		case '\r':
			continue
		case '\n':
			y++
			x = 0
			continue
		}
		cells[image.Point{X: x, Y: y}] = Cell{
			Rune:  r,
			Style: textbox.Style,
		}
		x++
		if textbox.Wrap && x > textWidth {
			x = 0
			y++
		}
		if y >= textHeight {
			break
		}
	}

	border := NewBorder()
	border.SetX(textbox.x)
	border.SetY(textbox.y)
	border.SetWidth(textbox.width)
	border.SetHeight(textbox.height)
	border.SetText(textbox.title)
	border.SetStyle(
		Style{
			BackgroundColor: textbox.Style.BackgroundColor,
			ForegroundColor: textbox.BorderColor,
		},
	)

	return append([]Tile{
		Tile{
			Rectangle: image.Rectangle{
				Min: image.Point{X: textbox.x + 1, Y: textbox.y + 1},
				Max: image.Point{X: textbox.x + 1 + textbox.width - 2, Y: textbox.y + 1 + textbox.height - 2},
			},
			Cells: cells,
		},
	},
		border.GetTiles(gui)...,
	)
}
