package gobless

import "image"

type TextBox struct {
	Width       int
	Height      int
	X           int
	Y           int
	Text        string
	Wrap        bool
	Border      bool
	Style       Style
	BorderColor Color
}

func NewTextBox() *TextBox {
	return &TextBox{
		Style:       DefaultStyle,
		BorderColor: NewColor(255, 255, 255),
	}
}

func (textbox *TextBox) GetTiles() []Tile {

	cells := map[image.Point]Cell{}

	textArea := image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: textbox.Width, Y: textbox.Height},
	}

	if textbox.Border {
		// keep text off borders
		textArea.Min.X++
		textArea.Min.Y++
		textArea.Max.X--
		textArea.Max.Y--

	}

	for x := 0; x < textbox.Width; x++ {
		for y := 0; y < textbox.Height; y++ {

			cell := Cell{
				Rune:  ' ',
				Style: textbox.Style,
			}
			if textbox.Border {
				if y == 0 {
					switch x {
					case 0:
						cell.Rune = '┌'
					case textArea.Max.X:
						cell.Rune = '┐'
					default:
						cell.Rune = '─'
					}
					cell.Style.ForegroundColor = textbox.BorderColor

				} else if y == textArea.Max.Y {
					switch x {
					case 0:
						cell.Rune = '└'
					case textArea.Max.X:
						cell.Rune = '┘'
					default:
						cell.Rune = '─'
					}
					cell.Style.ForegroundColor = textbox.BorderColor
				} else if x == 0 || x == textArea.Max.X {
					cell.Rune = '│'
					cell.Style.ForegroundColor = textbox.BorderColor
				}

			}
			cells[image.Point{X: x, Y: y}] = cell
		}
	}

	x, y := textArea.Min.X, textArea.Min.Y

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
				Min: image.Point{X: textbox.X, Y: textbox.Y},
				Max: image.Point{X: textbox.X + textbox.Width, Y: textbox.Y + textbox.Height},
			},
			Cells: cells,
		},
	}
}
