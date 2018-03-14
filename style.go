package gobless

import "github.com/gdamore/tcell"

type Style struct {
	ForegroundColor Color
	BackgroundColor Color
}

var DefaultStyle = Style{
	ForegroundColor: ColorWhite,
	BackgroundColor: ColorBlack,
}

func NewStyle(backgroundColor Color, foregroundColor Color) Style {
	return Style{
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
	}
}

func (style *Style) toTCell() tcell.Style {
	return tcell.StyleDefault.Foreground(style.ForegroundColor.toTCell()).Background(style.BackgroundColor.toTCell())
}
