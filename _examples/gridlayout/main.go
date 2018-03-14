package main

import "github.com/liamg/gobless"

func main() {
	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	helloTextbox := gobless.NewTextBox()
	helloTextbox.SetText(`Hello World!`)
	helloTextbox.SetBorderColor(gobless.ColorGreen)
	helloTextbox.SetTitle("Message")

	quitTextbox := gobless.NewTextBox()
	quitTextbox.SetText(`Press Ctrl-q to exit.`)
	quitTextbox.SetBorderColor(gobless.ColorRed)

	chart := gobless.NewBarChart()
	chart.SetTitle("Traffic")
	chart.SetBar("EU", 60)
	chart.SetBar("NA", 72)
	chart.SetBar("SA", 37)

	row := gobless.NewRow(
		gobless.NewColumn(
			gobless.ColumnSizeTwoThirds,
			helloTextbox,
		),
		gobless.NewColumn(
			gobless.ColumnSizeOneThird,
			chart,
		),
	)

	lowerRow := gobless.NewRow(
		gobless.NewColumn(
			gobless.ColumnSizeFull,
			quitTextbox,
		),
	)

	gui.Render(row, lowerRow)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}
