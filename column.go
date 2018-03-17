package gobless

type Column struct {
	components     []Component
	width          int
	height         int
	size           GridSize
	x              int
	y              int
	firstRowHeight int
	rowHeight      int
}

type GridSize int

const (
	GridSizeFull          GridSize = 12
	GridSizeFiveSixths    GridSize = 10
	GridSizeThreeQuarters GridSize = 9
	GridSizeTwoThirds     GridSize = 8
	GridSizeHalf          GridSize = 6
	GridSizeOneThird      GridSize = 4
	GridSizeOneQuarter    GridSize = 3
	GridSizeOneSixth      GridSize = 2
)

func NewColumn(size GridSize, components ...Component) *Column {
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

	for i, component := range col.components {

		component.SetX(col.x)
		component.SetY(col.y + yOffset)
		component.SetWidth(col.width)
		if i == 0 {
			component.SetHeight(col.firstRowHeight)
			yOffset += col.firstRowHeight
		} else {
			component.SetHeight(col.rowHeight)
			yOffset += col.rowHeight
		}

		tiles = append(tiles, component.GetTiles(gui)...)
	}

	return tiles
}
