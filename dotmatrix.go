package gobless

import (
	"image"
	"math"
)

type DotMatrix struct {
	width  int
	height int
	x      int
	y      int
	dots   map[int]map[int]*Color
}

func NewDotMatrix() *DotMatrix {
	return &DotMatrix{
		dots: map[int]map[int]*Color{},
	}
}

func (screen *DotMatrix) SetX(x int) {
	screen.x = x
}
func (screen *DotMatrix) SetY(y int) {
	screen.y = y
}
func (screen *DotMatrix) SetWidth(w int) {
	screen.width = w
}
func (screen *DotMatrix) SetHeight(h int) {
	screen.height = h
}

func (screen *DotMatrix) GetDotWidth() int {
	return screen.width * 2
}

func (screen *DotMatrix) GetDotHeight() int {
	return screen.height * 2
}

func (screen *DotMatrix) On(dotX int, dotY int, color Color) {
	screen.set(dotX, dotY, &color)
}

func (screen *DotMatrix) Off(dotX int, dotY int) {
	screen.set(dotX, dotY, nil)
}

func (screen *DotMatrix) set(dotX int, dotY int, color *Color) {
	if dotX < 0 || dotX >= screen.GetDotWidth() {
		return
	}
	if dotY < 0 || dotY >= screen.GetDotHeight() {
		return
	}
	if _, ok := screen.dots[dotX]; !ok {
		screen.dots[dotX] = map[int]*Color{}
	}
	screen.dots[dotX][dotY] = color
}

func (screen *DotMatrix) get(dotX int, dotY int) *Color {
	if _, ok := screen.dots[dotX]; !ok {
		return nil
	}
	if v, ok := screen.dots[dotX][dotY]; ok {
		return v
	}
	return nil
}

func (screen *DotMatrix) GetTiles(gui *GUI) []Tile {

	tile := NewTile(screen.x, screen.y, screen.x+screen.width-1, screen.y+screen.height-1)

	bg := DefaultStyle.BackgroundColor

	for x := 0; x < screen.width; x++ {
		for y := 0; y < screen.width; y++ {

			var color *Color

			topLeft := screen.get(x*2, y*2)
			topRight := screen.get((x*2)+1, y*2)
			bottomLeft := screen.get(x*2, (y*2)+1)
			bottomRight := screen.get((x*2)+1, (y*2)+1)
			if topLeft != nil {
				color = topLeft
			} else if topRight != nil {
				color = topRight
			} else if bottomLeft != nil {
				color = bottomLeft
			} else if bottomRight != nil {
				color = bottomRight
			}

			if color != nil {
				r := screen.getRune(
					screen.get(x*2, y*2) != nil,
					screen.get((x*2)+1, y*2) != nil,
					screen.get(x*2, (y*2)+1) != nil,
					screen.get((x*2)+1, (y*2)+1) != nil,
				)
				tile.SetCell(image.Point{X: x, Y: y}, NewCell(r, NewStyle(bg, *color)))
			} else {
				tile.SetCell(image.Point{X: x, Y: y}, NewCell(' ', DefaultStyle))
			}

		}
	}

	return []Tile{tile}
}

func (screen *DotMatrix) getRune(topLeft bool, topRight bool, bottomLeft bool, bottomRight bool) rune {

	switch true {
	case topLeft && topRight && bottomLeft && bottomRight:
		return '█'
	case topLeft && topRight && bottomLeft && !bottomRight:
		return '▛'
	case topLeft && topRight && !bottomLeft && bottomRight:
		return '▜'
	case topLeft && topRight && !bottomLeft && !bottomRight:
		return '▀'
	case topLeft && !topRight && bottomLeft && bottomRight:
		return '▙'
	case topLeft && !topRight && bottomLeft && !bottomRight:
		return '▌'
	case topLeft && !topRight && !bottomLeft && bottomRight:
		return '▚'
	case topLeft && !topRight && !bottomLeft && !bottomRight:
		return '▘'
	case !topLeft && topRight && bottomLeft && bottomRight:
		return '▟'
	case !topLeft && topRight && bottomLeft && !bottomRight:
		return '▞'
	case !topLeft && topRight && !bottomLeft && bottomRight:
		return '▐'
	case !topLeft && topRight && !bottomLeft && !bottomRight:
		return '▝'
	case !topLeft && !topRight && bottomLeft && bottomRight:
		return '▄'
	case !topLeft && !topRight && bottomLeft && !bottomRight:
		return '▖'
	case !topLeft && !topRight && !bottomLeft && bottomRight:
		return '▗'
	case !topLeft && !topRight && !bottomLeft && !bottomRight:
		return ' '
	default:
		return 'E'
	}
}

func (dotMatrix *DotMatrix) Line(x0 int, y0 int, x1 int, y1 int, color Color) {

	if math.Abs(float64(y1-y0)) > math.Abs(float64(x1-x0)) {

		// x = my + c
		m := float64(x1-x0) / float64(y1-y0)
		c := int(math.Round(float64(x0) - (m * float64(y0))))

		inc := 1
		if y0 > y1 {
			inc = -1
		}

		for y := y0; y != y1; y += inc {
			dotMatrix.On(int(math.Round((m*float64(y))+float64(c))), y, color)
		}

	} else {

		// y = mx + c
		m := float64(y1-y0) / float64(x1-x0)
		c := int(math.Round(float64(y0) - (m * float64(x0))))

		inc := 1
		if x0 > x1 {
			inc = -1
		}

		for x := x0; x != x1; x += inc {
			dotMatrix.On(x, int(math.Round((m*float64(x))+float64(c))), color)
		}
	}
}

func (dotMatrix *DotMatrix) Circle(x0 int, y0 int, radius int, color Color) {

	x := radius - 1
	y := 0
	dx := 1
	dy := 1
	err := dx - (radius << 1)

	for x >= y {

		dotMatrix.On(x0+x, y0+y, color)
		dotMatrix.On(x0+y, y0+x, color)
		dotMatrix.On(x0-y, y0+x, color)
		dotMatrix.On(x0-x, y0+y, color)
		dotMatrix.On(x0-x, y0-y, color)
		dotMatrix.On(x0-y, y0-x, color)
		dotMatrix.On(x0+y, y0-x, color)
		dotMatrix.On(x0+x, y0-y, color)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}

		if err > 0 {
			x--
			dx += 2
			err += dx - (radius << 1)
		}
	}
}
