package gobless

import "github.com/gdamore/tcell"

type Style struct {
	ForegroundColor Color
	BackgroundColor Color
}

var DefaultStyle = Style{
	ForegroundColor: NewColor(0xff, 0xff, 0xff),
	BackgroundColor: NewColor(0x00, 0x00, 0x00),
}

type Color int32

func NewColor(r uint32, g uint32, b uint32) Color {
	return Color((1 << 24) | ((r & 0xff) << 16) | ((g & 0xff) << 8) | (b & 0xff))
}

func (color *Color) toTCell() tcell.Color {
	return tcell.Color(*color)
}

func (style *Style) toTCell() tcell.Style {
	return tcell.StyleDefault.Foreground(style.ForegroundColor.toTCell()).Background(style.BackgroundColor.toTCell())
}
