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

	textbox := gobless.NewTextBox()
	textbox.SetX(2)
	textbox.SetY(2)
	textbox.SetWidth(24)
	textbox.SetHeight(4)
	textbox.Border = true
	textbox.BorderColor = gobless.NewColor(255, 0, 0)
	textbox.Wrap = true
	textbox.Text = `This line should wrap because it is long...`
	textbox.Style = gobless.Style{
		ForegroundColor: gobless.NewColor(192, 0, 0),
		BackgroundColor: gobless.NewColor(0, 0, 0),
	}

	gui.Render(textbox)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()

}
