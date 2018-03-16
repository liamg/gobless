package gobless

import (
	"math"
)

type Row struct {
	columns []*Column
	width   int
	height  int
	x       int
	y       int
}

func NewRow(columns ...*Column) *Row {
	return &Row{
		columns: columns,
	}
}

func (row *Row) SetWidth(w int) {
	row.width = w
}
func (row *Row) SetHeight(h int) {
	row.height = h
}

func (row *Row) SetX(x int) {
	row.x = x
}
func (row *Row) SetY(y int) {
	row.y = y
}
func (row *Row) GetTiles(gui *GUI) []Tile {

	tiles := []Tile{}

	colWidth := float64(row.width) / 12

	flSum := 0.0

	colOffset := row.x

	for _, col := range row.columns {

		flWidth := colWidth * float64(col.size)
		col.width = int(math.Floor(flWidth))
		flSum += (flWidth - float64(col.width))

		col.x = colOffset
		col.y = row.y
		colOffset += col.width

		if flSum >= 0.999 { // avoid floating point fun :/
			flSum--
			col.width = col.width + 1

		}

		col.height = row.height
		col.rowHeight = int(math.Floor(float64(row.height) / float64(len(col.components))))
		col.firstRowHeight = col.rowHeight + (row.height - (len(col.components) * col.rowHeight))

		tiles = append(tiles, col.GetTiles(gui)...)
	}

	return tiles
}
