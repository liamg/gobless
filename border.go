package gobless

import (
	"image"
)

type Border struct {
	width  int
	height int
	x      int
	y      int
	text   string
	style  Style
}

func NewBorder() *Border {
	style := DefaultStyle
	style.ForegroundColor = NewColor(0, 255, 255)
	return &Border{
		style: style,
	}
}

func (component *Border) SetX(x int) {
	component.x = x
}
func (component *Border) SetY(y int) {
	component.y = y
}
func (component *Border) SetWidth(w int) {
	component.width = w
}
func (component *Border) SetHeight(h int) {
	component.height = h
}
func (component *Border) SetText(text string) {
	component.text = text
}
func (component *Border) SetStyle(style Style) {
	component.style = style
}

func (component *Border) GetTiles(gui *GUI) []Tile {

	topLeft := '┌'
	topRight := '┐'
	bottomLeft := '└'
	vertical := '│'
	horizontal := '─'
	bottomRight := '┘'

	topTile := Tile{
		Rectangle: image.Rectangle{
			Min: image.Point{
				X: component.x,
				Y: component.y,
			},
			Max: image.Point{
				X: component.x + component.width - 1,
				Y: component.y,
			},
		},
		Cells: map[image.Point]Cell{},
	}
	bottomTile := Tile{
		Rectangle: image.Rectangle{
			Min: image.Point{
				X: component.x,
				Y: component.y + component.height - 1,
			},
			Max: image.Point{
				X: component.x + component.width - 1,
				Y: component.y + component.height - 1,
			},
		},
		Cells: map[image.Point]Cell{},
	}
	leftTile := Tile{
		Rectangle: image.Rectangle{
			Min: image.Point{
				X: component.x,
				Y: component.y,
			},
			Max: image.Point{
				X: component.x,
				Y: component.y + component.height - 1,
			},
		},
		Cells: map[image.Point]Cell{},
	}
	rightTile := Tile{
		Rectangle: image.Rectangle{
			Min: image.Point{
				X: component.x + component.width - 1,
				Y: component.y,
			},
			Max: image.Point{
				X: component.x + component.width - 1,
				Y: component.y + component.height - 1,
			},
		},
		Cells: map[image.Point]Cell{},
	}

	textOffsetFromCorner := 2

	for x := 0; x < component.width; x++ {

		borderRune := horizontal

		if x == 0 {
			borderRune = bottomLeft
		} else if x == component.width-1 {
			borderRune = bottomRight
		}

		bottomTile.Cells[image.Point{X: x, Y: 0}] = Cell{
			Rune:  borderRune,
			Style: component.style,
		}

		if x == 0 {
			borderRune = topLeft
		} else if x == component.width-1 {
			borderRune = topRight
		} else if x >= textOffsetFromCorner && x-textOffsetFromCorner < len(component.text) {
			borderRune = []rune(component.text)[x-textOffsetFromCorner]
		}

		topTile.Cells[image.Point{X: x, Y: 0}] = Cell{
			Rune:  borderRune,
			Style: component.style,
		}
	}

	for y := 1; y < component.height-1; y++ {
		leftTile.Cells[image.Point{X: 0, Y: y}] = Cell{
			Rune:  vertical,
			Style: component.style,
		}
		rightTile.Cells[image.Point{X: 0, Y: y}] = Cell{
			Rune:  vertical,
			Style: component.style,
		}
	}

	return []Tile{topTile, bottomTile, leftTile, rightTile}
}
