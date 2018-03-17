package main

import (
	"math/rand"
	"time"

	"github.com/liamg/gobless"
)

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
	chart.SetBar("Europe", 90)
	chart.SetBar("US", 72)
	chart.SetBar("Asia", 12)
	chart.SetYScale(100)

	row := gobless.NewRow(
		gobless.GridSizeHalf,
		gobless.NewColumn(
			gobless.GridSizeTwoThirds,
			helloTextbox,
		),
		gobless.NewColumn(
			gobless.GridSizeOneThird,
			chart,
		),
	)

	lowerRow := gobless.NewRow(
		gobless.GridSizeHalf,
		gobless.NewColumn(
			gobless.GridSizeFull,
			quitTextbox,
		),
	)

	gui.Render(row, lowerRow)

	quitChan := make(chan bool)

	rand.Seed(time.Now().UnixNano())

	go func() {

		for {
			select {
			case <-time.After(time.Second):
				chart.SetBar("Europe", rand.Intn(30)+70)
				chart.SetBar("US", rand.Intn(20)+50)
				chart.SetBar("Asia", rand.Intn(10)+5)
				gui.Render(row, lowerRow)
			case <-quitChan:
				break
			}
		}

	}()

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		quitChan <- true
		gui.Close()
	})

	gui.HandleResize(func(event gobless.ResizeEvent) {
		gui.Render(row, lowerRow)
	})

	gui.Loop()
}
