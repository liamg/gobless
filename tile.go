package gobless

import "image"

type Tile struct {
	Rectangle image.Rectangle
	Cells     map[image.Point]Cell
}
