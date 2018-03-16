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
	quitTextbox.SetBorderColor(gobless.ColorDarkRed)

	chart := gobless.NewBarChart()
	chart.SetTitle("Traffic")
	chart.SetYScale(100)
	chart.SetBar("EU", 60)
	chart.SetBar("NA", 72)
	chart.SetBar("SA", 37)

	chart2 := gobless.NewBarChart()
	chart2.SetTitle("Wasp Count")
	chart2.SetYScale(10000000)
	chart2.SetBarStyle(gobless.NewStyle(gobless.ColorRed, gobless.ColorWhite))
	chart2.SetBorderColor(gobless.ColorRed)
	chart2.SetBar("EU", 3500000)
	chart2.SetBar("NA", 1100000)
	chart2.SetBar("SA", 9400000)

	rows := []gobless.Component{
		gobless.NewRow(
			gobless.NewColumn(
				gobless.ColumnSizeTwoThirds,
				helloTextbox,
			),
			gobless.NewColumn(
				gobless.ColumnSizeOneThird,
				gobless.NewRow(
					gobless.NewColumn(
						gobless.ColumnSizeFull,
						chart,
						chart2,
					),
				),
			),
		), gobless.NewRow(
			gobless.NewColumn(
				gobless.ColumnSizeFull,
				quitTextbox,
			),
		),
	}

	gui.Render(rows...)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.HandleResize(func(event gobless.ResizeEvent) {
		gui.Render(rows...)
	})

	gui.Loop()
}
