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
}

func NewTextBox() *TextBox {
	return &TextBox{
		Style:       DefaultStyle,
		BorderColor: NewColor(0, 255, 255),
	}
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

	textArea := image.Rectangle{
		Min: image.Point{X: 1, Y: 1},
		Max: image.Point{X: textbox.width - 2, Y: textbox.height - 2},
	}

	for x := 0; x < textbox.width; x++ {
		for y := 0; y < textbox.height; y++ {

			cell := Cell{
				Rune:  ' ',
				Style: textbox.Style,
			}

			if y == 0 {
				switch x {
				case 0:
					cell.Rune = '┌'
				case textbox.width - 1:
					cell.Rune = '┐'
				default:
					cell.Rune = '─'
				}
				cell.Style.ForegroundColor = textbox.BorderColor

			} else if y == textbox.height-1 {
				switch x {
				case 0:
					cell.Rune = '└'
				case textbox.width - 1:
					cell.Rune = '┘'
				default:
					cell.Rune = '─'
				}
				cell.Style.ForegroundColor = textbox.BorderColor
			} else if x == 0 || x == textbox.width-1 {
				cell.Rune = '│'
				cell.Style.ForegroundColor = textbox.BorderColor
			}

			cells[image.Point{X: x, Y: y}] = cell
		}
	}

	x, y := textArea.Min.X, textArea.Min.Y

	//textbox.Text = fmt.Sprintf("%d,%d %d,%d", textbox.x, textbox.y, textbox.width, textbox.height)

	for _, r := range []rune(textbox.Text) {
		if !textbox.Wrap && x >= textArea.Max.X {
			continue
		}
		switch r {
		case '\r':
			continue
		case '\n':
			y++
			x = textArea.Min.X
			continue
		}
		cells[image.Point{X: x, Y: y}] = Cell{
			Rune:  r,
			Style: textbox.Style,
		}
		x++
		if textbox.Wrap && x >= textArea.Max.X {
			x = textArea.Min.X
			y++
		}
		if y >= textArea.Max.Y {
			break
		}
	}

	return []Tile{
		Tile{
			Rectangle: image.Rectangle{
				Min: image.Point{X: textbox.x, Y: textbox.y},
				Max: image.Point{X: textbox.x + textbox.width - 1, Y: textbox.y + textbox.height - 1},
			},
			Cells: cells,
		},
	}
}
