package main

import "github.com/liamg/gobless"

func main() {
	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	helloTextbox := gobless.NewTextBox()
	helloTextbox.X = 10
	helloTextbox.Y = 3
	helloTextbox.Width = 14
	helloTextbox.Height = 3
	helloTextbox.Text = `Hello World!`
	helloTextbox.Border = true

	quitTextbox := gobless.NewTextBox()
	quitTextbox.Width = 20
	quitTextbox.Height = 1
	quitTextbox.Text = `Press Ctrl-q to exit.`

	gui.Render(helloTextbox, quitTextbox)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}
