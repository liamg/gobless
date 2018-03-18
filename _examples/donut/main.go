package main

import "github.com/liamg/gobless"

func main() {
	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	donut := gobless.NewDonut()
	donut.SetY(4)
	donut.SetX(1)
	donut.SetWidth(18)
	donut.SetHeight(18)

	donut.SetLabel("EVIL")
	donut.SetPercent(80)

	gui.Render(donut)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}
