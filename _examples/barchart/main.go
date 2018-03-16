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

	chart := gobless.NewBarChart()
	chart.SetTitle("Traffic")
	chart.SetX(2)
	chart.SetY(1)
	chart.SetWidth(40)
	chart.SetHeight(15)
	chart.SetYScale(100)
	chart.SetBar("Europe", 60)
	chart.SetBar("US", 72)
	chart.SetBar("Asia", 37)

	quitChan := make(chan bool)

	rand.Seed(time.Now().UnixNano())

	go func() {

		for {

			select {
			case <-time.After(time.Second):
				chart.SetBar("Europe", rand.Intn(30)+70)
				chart.SetBar("US", rand.Intn(20)+50)
				chart.SetBar("Asia", rand.Intn(10)+5)
				gui.Render(chart)
			case <-quitChan:
				break
			}
		}

	}()

	gui.Render(chart)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		quitChan <- true
		gui.Close()
	})

	gui.HandleResize(func(event gobless.ResizeEvent) {
		gui.Render(chart)
	})

	gui.Loop()
}
