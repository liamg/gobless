# gobless
Build beautiful terminal dashboards and GUIs in Golang. 

The name was intended to be pronounced _go-bless_ due to the inspiration from [blessed-contrib](https://github.com/yaronn/blessed-contrib), but _gob-less_ is definitely funnier, so we'll go with that.

## Terminology

- A *Cell* refers to an individual terminal cell, in which a single character can be drawn.
- A *Tile* refers to a rectangular grouping of cells.
- A *Component* is an entity which provides a tile - or tiles - to the GUI to render. A bar chart, for example.

## Get Started

The below code shows a simple "Hello World" application. Additionally, there are several [examples](examples/) provided.

```
package main

import "github.com/liamg/gobless"

func main() {
    gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	helloTextbox := gobless.NewTextBox()
	helloTextbox.X = 0
	helloTextbox.Y = 0
	helloTextbox.Width = 14
	helloTextbox.Height = 3
	helloTextbox.Text = `Hello World!`
    helloTextbox.Border = true

    quitTextbox := gobless.NewTextBox()
	quitTextbox.Width = 20
	quitTextbox.Height = 1
	quitTextbox.Text = `Press Ctrl-q to exit.`

	gui.Render(textbox)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}

```