package gobless

type Cell struct {
	Rune  rune
	Style Style
}

func NewCell(c rune, style Style) Cell {
	return Cell{
		Rune:  c,
		Style: style,
	}
}
