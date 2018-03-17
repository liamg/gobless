package gobless

import (
	"image"
)

type TextBox struct {
	width       int
	height      int
	x           int
	y           int
	text        string
	wrap        bool
	style       Style
	borderColor Color
	title       string
}

func NewTextBox() *TextBox {
	return &TextBox{
		style:       DefaultStyle,
		borderColor: ColorDarkCyan,
	}
}

func (textbox *TextBox) SetTitle(title string) {
	textbox.title = title
}
func (textbox *TextBox) SetBorderColor(color Color) {
	textbox.borderColor = color
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

func (component *TextBox) SetTextWrap(enabled bool) {
	component.wrap = enabled
}
func (component *TextBox) SetText(text string) {
	component.text = text
}
func (component *TextBox) SetStyle(style Style) {
	component.style = style
}

func (textbox *TextBox) GetTiles(gui *GUI) []Tile {

	cells := map[image.Point]Cell{}

	x, y := 0, 0

	textWidth := textbox.width - 3 // remove border dimensions
	textHeight := textbox.height - 2

	for _, r := range []rune(textbox.text) {
		if !textbox.wrap && x >= textWidth {
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
			Style: textbox.style,
		}
		x++
		if textbox.wrap && x > textWidth {
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
	border.SetStyle(NewStyle(textbox.style.BackgroundColor, textbox.borderColor))

	return append([]Tile{
		Tile{
			Rectangle: image.Rectangle{
				Min: image.Point{X: textbox.x + 1, Y: textbox.y + 1},
				Max: image.Point{X: textbox.x + 1 + textbox.width - 3, Y: textbox.y + 1 + textbox.height - 3},
			},
			Cells: cells,
		},
	},
		border.GetTiles(gui)...,
	)
}
