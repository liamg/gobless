package gobless

type Donut struct {
	width  int
	height int
	x      int
	y      int
	style  Style
}

func NewDonut() *Donut {
	return &Donut{
		style: DefaultStyle,
	}
}

func (donut *Donut) SetX(x int) {
	donut.x = x
}
func (donut *Donut) SetY(y int) {
	donut.y = y
}
func (donut *Donut) SetWidth(w int) {
	donut.width = w
}
func (donut *Donut) SetHeight(h int) {
	donut.height = h
}
func (donut *Donut) SetStyle(style Style) {
	donut.style = style
}

func (donut *Donut) GetTiles(gui *GUI) []Tile {

	tile := NewTile(donut.x, donut.y, donut.x+donut.width-1, donut.y+donut.height-1)

	return []Tile{tile}
}
