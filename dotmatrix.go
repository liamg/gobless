package gobless

import (
	"image"
)

type DotMatrix struct {
	width  int
	height int
	x      int
	y      int
	style  Style
	dots   map[int]map[int]bool
}

func NewDotMatrix() *DotMatrix {
	return &DotMatrix{
		style: DefaultStyle,
		dots:  map[int]map[int]bool{},
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
func (screen *DotMatrix) SetStyle(style Style) {
	screen.style = style
}

func (screen *DotMatrix) GetDotWidth() int {
	return screen.width * 2
}

func (screen *DotMatrix) GetDotHeight() int {
	return screen.height * 2
}

func (screen *DotMatrix) On(dotX int, dotY int) {
	screen.set(dotX, dotY, true)
}

func (screen *DotMatrix) Off(dotX int, dotY int) {
	screen.set(dotX, dotY, false)
}

func (screen *DotMatrix) set(dotX int, dotY int, val bool) {
	if dotX < 0 || dotX >= screen.GetDotWidth() {
		return
	}
	if dotY < 0 || dotY >= screen.GetDotHeight() {
		return
	}
	if _, ok := screen.dots[dotX]; !ok {
		screen.dots[dotX] = map[int]bool{}
	}
	screen.dots[dotX][dotY] = val
}

func (screen *DotMatrix) get(dotX int, dotY int) bool {
	if _, ok := screen.dots[dotX]; !ok {
		return false
	}
	if v, ok := screen.dots[dotX][dotY]; ok {
		return v
	}
	return false
}

func (screen *DotMatrix) GetTiles(gui *GUI) []Tile {

	tile := NewTile(screen.x, screen.y, screen.x+screen.width-1, screen.y+screen.height-1)

	for x := 0; x < screen.width; x++ {
		for y := 0; y < screen.width; y++ {
			tile.SetCell(image.Point{X: x, Y: y}, NewCell(screen.getRune(
				screen.get(x*2, y*2),
				screen.get((x*2)+1, y*2),
				screen.get(x*2, (y*2)+1),
				screen.get((x*2)+1, (y*2)+1),
			), screen.style))
		}
	}

	return []Tile{tile}
}

func (screen *DotMatrix) getRune(topLeft bool, topRight bool, bottomLeft bool, bottomRight bool) rune {

	switch true {
	case topLeft && topRight && bottomLeft && bottomRight:
		return '█'
	case !topLeft && topRight && bottomLeft && bottomRight:
		return '▟'
	case topLeft && !topRight && bottomLeft && bottomRight:
		return '▙'
	case topLeft && !topRight && !bottomLeft && bottomRight:
		return '▚'
	case topLeft && topRight && bottomLeft && !bottomRight:
		return '▛'
	case !topLeft && topRight && bottomLeft && bottomRight:
		return '▜'
	case !topLeft && topRight && bottomLeft && !bottomRight:
		return '▞'
	case !topLeft && topRight && !bottomLeft && !bottomRight:
		return '▝'
	case !topLeft && !topRight && bottomLeft && bottomRight:
		return '▄'
	case topLeft && topRight && !bottomLeft && !bottomRight:
		return '▀'
	case !topLeft && !topRight && bottomLeft && !bottomRight:
		return '▖'
	case !topLeft && !topRight && !bottomLeft && bottomRight:
		return '▗'
	case topLeft && !topRight && !bottomLeft && !bottomRight:
		return '▘'
	case topLeft && !topRight && bottomLeft && !bottomRight:
		return '▌'
	case !topLeft && topRight && !bottomLeft && bottomRight:
		return '▐'
	case !topLeft && !topRight && !bottomLeft && !bottomRight:
		return ' '
	default:
		return 'E'
	}
}
