package gobless

type Component interface {
	GetTiles(gui *GUI) []Tile
	SetX(x int)
	SetY(y int)
	SetWidth(w int)
	SetHeight(h int)
}
