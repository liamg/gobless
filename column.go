package gobless

type Column struct {
	components []Component
	width      int
	height     int
	size       ColumnSize
	x          int
	y          int
	rowHeight  int
}

type ColumnSize int

const (
	ColumnSizeFull          ColumnSize = 12
	ColumnSizeFiveSixths    ColumnSize = 10
	ColumnSizeThreeQuarters ColumnSize = 9
	ColumnSizeTwoThirds     ColumnSize = 8
	ColumnSizeHalf          ColumnSize = 6
	ColumnSizeThird         ColumnSize = 4
	ColumnSizeQuarter       ColumnSize = 3
	ColumnSizeOneSixth      ColumnSize = 2
)

func NewColumn(size ColumnSize, components ...Component) *Column {
	return &Column{
		components: components,
		size:       size,
		rowHeight:  10,
	}
}

func (col *Column) SetX(x int) {
	col.x = x
}
func (col *Column) SetY(y int) {
	col.y = y
}
func (col *Column) SetWidth(w int) {
	col.width = w
}
func (col *Column) SetHeight(h int) {
	col.height = h
}

func (col *Column) GetTiles(gui *GUI) []Tile {
	tiles := []Tile{}

	yOffset := 0

	for _, component := range col.components {
		component.SetX(col.x)
		component.SetY(col.y + yOffset)
		component.SetWidth(col.width)
		component.SetHeight(col.rowHeight)
		yOffset += col.rowHeight
		tiles = append(tiles, component.GetTiles(gui)...)
	}

	return tiles
}
