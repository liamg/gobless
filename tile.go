package gobless

import "image"

type Tile struct {
	Rectangle image.Rectangle
	Cells     map[image.Point]Cell
}

func NewTile(x1 int, y1 int, x2 int, y2 int) Tile {
	return Tile{
		Rectangle: image.Rect(x1, y1, x2, y2),
		Cells:     map[image.Point]Cell{},
	}
}

func (tile *Tile) SetCell(point image.Point, cell Cell) {
	tile.Cells[point] = cell
}

func (tile *Tile) SetCells(rectangle image.Rectangle, cell Cell) {
	for x := rectangle.Min.X; x <= rectangle.Max.X; x++ {
		for y := rectangle.Min.Y; y <= rectangle.Max.Y; y++ {
			tile.SetCell(image.Point{X: x, Y: y}, cell)
		}
	}
}
