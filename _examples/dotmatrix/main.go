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

	dotMatrix.Line(0, 0, 0, 10, gobless.ColorRed)

	quitTextbox := gobless.NewTextBox()
	quitTextbox.SetWidth(24)
	quitTextbox.SetHeight(3)
	quitTextbox.SetText(`Press Ctrl-q to exit.`)

	dotMatrix2 := gobless.NewDotMatrix()
	dotMatrix2.SetY(4)
	dotMatrix2.SetX(24)
	dotMatrix2.SetWidth(22)
	dotMatrix2.SetHeight(22)

	dotMatrix2.Circle(20, 20, 10, gobless.ColorRebeccaPurple)
	dotMatrix2.Circle(20, 20, 15, gobless.ColorLimeGreen)
	dotMatrix2.Circle(20, 20, 20, gobless.ColorMaroon)

	gui.Render(dotMatrix, quitTextbox, dotMatrix2)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}
