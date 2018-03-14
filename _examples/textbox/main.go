package main

import "github.com/liamg/gobless"

func main() {
	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	helloTextbox := gobless.NewTextBox()
	helloTextbox.SetX(10)
	helloTextbox.SetY(3)
	helloTextbox.SetWidth(14)
	helloTextbox.SetHeight(3)
	helloTextbox.SetText(`Hello World!`)

	quitTextbox := gobless.NewTextBox()
	quitTextbox.SetWidth(22)
	quitTextbox.SetHeight(3)
	quitTextbox.SetText(`Press Ctrl-q to exit.`)

	gui.Render(helloTextbox, quitTextbox)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}
