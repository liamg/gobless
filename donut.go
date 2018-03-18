package gobless

import "math"

type Donut struct {
	width  int
	height int
	x      int
	y      int
	style  Style
	pc     int
	label  string
}

func NewDonut() *Donut {
	return &Donut{
		style: DefaultStyle,
	}
}

func (donut *Donut) SetPercent(pc int) {
	donut.pc = pc
}
func (donut *Donut) SetLabel(label string) {
	donut.label = label
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

	dm := NewDotMatrix()
	dm.SetWidth(donut.width)
	dm.SetHeight(donut.height)

	cX := donut.width
	cY := donut.height

	finalAngle := float64(donut.pc) * 360.0

	maxRadius := cX
	if maxRadius > cY {
		maxRadius = cY
	}

	for radius := 1; radius <= maxRadius; radius++ {

		inc := 2.0 / float64(radius)

		c2 := float64(radius * radius)

		for deg := 0.0; deg < finalAngle; deg += inc {
			subRad := math.Mod(deg, 90.0) * (math.Pi / 180)
			y := math.Sin(subRad) * float64(radius)
			x := math.Sqrt(c2 - (y * y))

			switch true {
			case deg > 270:
				x = x * -1
				y = y * -1
			case deg > 180:
				x = x * -1
			case deg > 90:
			case deg > 0:
				y = y * -1
			}

			dm.Line(cX, cY, cX+int(math.Round(x)), cY+int(math.Round(y)), ColorRed)
		}
	}

	return dm.GetTiles(gui)
}
