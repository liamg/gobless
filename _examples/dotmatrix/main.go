package main

import "github.com/liamg/gobless"

func main() {
	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	dotMatrix := gobless.NewDotMatrix()
	dotMatrix.SetY(4)
	dotMatrix.SetX(1)
	dotMatrix.SetWidth(22)
	dotMatrix.SetHeight(22)
	dotMatrix.SetStyle(
		gobless.NewStyle(
			gobless.DefaultStyle.BackgroundColor,
			gobless.ColorLightSlateGray,
		),
	)

	xSpeed, ySpeed := 1, 0
	xMax, yMax := 40, 40
	xMin, yMin := 0, 2

	for x, y := 0, 0; true; x, y = x+xSpeed, y+ySpeed {
		dotMatrix.On(x, y)

		switch true {
		case xSpeed > 0 && x == xMax:
			xMax -= 2
			xSpeed = 0
			ySpeed = 1
		case ySpeed > 0 && y == yMax:
			yMax -= 2
			ySpeed = 0
			xSpeed = -1
		case xSpeed < 0 && x == xMin:
			xMin += 2
			xSpeed = 0
			ySpeed = -1
		case ySpeed < 0 && y == yMin:
			yMin += 2
			xSpeed = 1
			ySpeed = 0
		}

		if xMin > xMax || yMin > yMax {
			break
		}

	}

	quitTextbox := gobless.NewTextBox()
	quitTextbox.SetWidth(22)
	quitTextbox.SetHeight(3)
	quitTextbox.SetText(`Press Ctrl-q to exit.`)

	gui.Render(dotMatrix, quitTextbox)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}
