package main

import "github.com/liamg/gobless"

func main() {
	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	chart := gobless.NewBarChart()
	chart.SetTitle("Traffic")
	chart.SetWidth(40)
	chart.SetHeight(15)
	chart.SetBar("EU", 60)
	chart.SetBar("NA", 72)
	chart.SetBar("SA", 37)

	gui.Render(chart)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}
