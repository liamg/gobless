package main

import "github.com/liamg/gobless"

func main() {
	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	helloTextbox := gobless.NewTextBox()
	helloTextbox.SetText(`Bee and wasp stings can be painful and in some people they can trigger a life-threatening reaction. 
		
Symptoms include:

 - Wheezing or difficulty breathing
 - Severe swelling of the face, throat, or lips
 - Nausea, vomiting, or diarrhea
 - Stomach cramps
 - Itching or hives (in places other than the site of the sting)
 - Fast heart rate
 - Sudden death`)
	helloTextbox.SetBorderColor(gobless.ColorGreen)
	helloTextbox.SetTitle("Information")
	helloTextbox.SetTextWrap(true)

	quitTextbox := gobless.NewTextBox()
	quitTextbox.SetText(`Press Ctrl-q to exit.`)
	quitTextbox.SetBorderColor(gobless.ColorDarkRed)

	chart := gobless.NewBarChart()
	chart.SetTitle("Diarrhoea")
	chart.SetYScale(100)
	chart.SetBar("EU", 60)
	chart.SetBar("NA", 72)
	chart.SetBar("SA", 37)
	chart.SetBarStyle(gobless.NewStyle(
		gobless.ColorOliveDrab,
		gobless.DefaultStyle.ForegroundColor,
	))

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
			gobless.GridSizeThreeQuarters,
			gobless.NewColumn(
				gobless.GridSizeTwoThirds,
				helloTextbox,
			),
			gobless.NewColumn(
				gobless.GridSizeOneThird,
				gobless.NewRow(
					gobless.GridSizeFull,
					gobless.NewColumn(
						gobless.GridSizeFull,
						chart,
						chart2,
					),
				),
			),
		), gobless.NewRow(
			gobless.GridSizeOneQuarter,
			gobless.NewColumn(
				gobless.GridSizeFull,
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
