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
	helloTextbox.SetWidth(64)
	helloTextbox.SetHeight(10)
	helloTextbox.SetTextWrap(true)
	helloTextbox.SetText(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc gravida vulputate augue, ut lobortis neque semper at. In hac habitasse platea dictumst. In tincidunt diam vitae bibendum posuere. Maecenas imperdiet nunc non ex dignissim rhoncus. Integer consectetur turpis sit amet fermentum eleifend. Fusce placerat urna in purus efficitur dictum ut at nibh. Cras bibendum arcu eget ligula posuere, a tristique dolor mattis. Vivamus quis lectus cursus, eleifend leo sit amet, pellentesque dolor.`)

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
