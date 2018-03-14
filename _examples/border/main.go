package main

import (
	"github.com/liamg/gobless"
)

func main() {

	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	border := gobless.NewBorder()
	border.SetWidth(8)
	border.SetHeight(4)
	border.SetX(15)
	border.SetY(5)

	gui.Render(border)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()

}
