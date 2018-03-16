package gobless

import (
	"image"
	"math"
)

type BarChart struct {
	width       int
	height      int
	x           int
	y           int
	wrap        bool
	style       Style
	barstyle    Style
	borderColor Color
	title       string
	bars        []barChartBar
	space       bool
	yScale      int
}

type barChartBar struct {
	Name  string
	Value int
}

func NewBarChart() *BarChart {
	return &BarChart{
		style:       DefaultStyle,
		barstyle:    NewStyle(ColorDarkBlue, ColorWhite),
		borderColor: ColorDarkCyan,
		bars:        []barChartBar{},
		space:       true,
	}
}

func (barchart *BarChart) SetBarSpacing(space bool) {
	barchart.space = space
}
func (barchart *BarChart) SetYScale(scale int) {
	barchart.yScale = scale
}
func (barchart *BarChart) SetTitle(title string) {
	barchart.title = title
}
func (barchart *BarChart) SetBorderColor(color Color) {
	barchart.borderColor = color
}

func (barchart *BarChart) SetBar(name string, value int) {

	for i, bar := range barchart.bars {
		if bar.Name == name {
			barchart.bars[i].Value = value
			return
		}
	}

	barchart.bars = append(barchart.bars, barChartBar{name, value})
}

func (barchart *BarChart) SetX(x int) {
	barchart.x = x
}
func (barchart *BarChart) SetY(y int) {
	barchart.y = y
}
func (barchart *BarChart) SetWidth(w int) {
	barchart.width = w
}
func (barchart *BarChart) SetHeight(h int) {
	barchart.height = h
}

func (component *BarChart) SetTextWrap(enabled bool) {
	component.wrap = enabled
}
func (component *BarChart) SetStyle(style Style) {
	component.style = style
}
func (component *BarChart) SetBarStyle(style Style) {
	component.barstyle = style
}

func (barchart *BarChart) GetTiles(gui *GUI) []Tile {

	tile := NewTile(barchart.x+1, barchart.y+1, barchart.x+barchart.width-2, barchart.y+barchart.height-2)

	maxRenderWidth := barchart.width - 2

	startX := 0

	// bar area = bar width + bar spacing

	widthPerBarAreaFloat := float64(maxRenderWidth) / float64(len(barchart.bars))
	widthPerBarArea := int(math.Floor(widthPerBarAreaFloat))
	widthPerBar := widthPerBarArea

	if barchart.space && widthPerBar > 1 {
		widthPerBar = widthPerBar - 1
	}

	leftoverChars := maxRenderWidth - (widthPerBarArea * len(barchart.bars))

	startX = int(leftoverChars / 2)

	maxBarHeight := barchart.height - 3

	maxValue := barchart.yScale

	if maxValue == 0 {
		for _, bar := range barchart.bars {
			if bar.Value > maxValue {
				maxValue = bar.Value
			}
		}
	}

	for index, bar := range barchart.bars {

		barHeight := int(math.Round(
			(float64(bar.Value) / float64(maxValue)) * float64(maxBarHeight),
		))

		if barHeight > maxBarHeight {
			barHeight = maxBarHeight
		} else if barHeight < 0 {
			barHeight = 0
		}

		barTopOffset := maxBarHeight - barHeight

		tile.SetCells(
			image.Rect(startX, 0, startX+widthPerBar-1, maxBarHeight-1),
			NewCell(' ', barchart.style),
		)

		if barHeight > 0 {
			tile.SetCells(
				image.Rect(startX, barTopOffset, startX+widthPerBar-1, maxBarHeight-1),
				NewCell(' ', barchart.barstyle),
			)

			label := shortenInt(bar.Value)

			for p := 0; p < len(label) && p < widthPerBar; p++ {
				tile.SetCell(image.Point{X: startX + p, Y: barTopOffset}, NewCell([]rune(label)[p], barchart.barstyle))
			}
		} else {
			label := shortenInt(bar.Value)

			for p := 0; p < len(label) && p < widthPerBar; p++ {
				tile.SetCell(image.Point{X: startX + p, Y: barTopOffset - 1}, NewCell([]rune(label)[p], barchart.style))
			}
		}

		for p := 0; p < widthPerBar; p++ {
			if p < len(bar.Name) {
				tile.SetCell(image.Point{X: startX + p, Y: maxBarHeight}, NewCell([]rune(bar.Name)[p], barchart.style))
			} else {
				tile.SetCell(image.Point{X: startX + p, Y: maxBarHeight}, NewCell(' ', barchart.style))
			}
		}

		startX += widthPerBarArea
		index++
	}

	border := NewBorder()
	border.SetX(barchart.x)
	border.SetY(barchart.y)
	border.SetWidth(barchart.width)
	border.SetHeight(barchart.height)
	border.SetText(barchart.title)
	border.SetStyle(NewStyle(barchart.style.BackgroundColor, barchart.borderColor))
	return append([]Tile{tile}, border.GetTiles(gui)...)
}
