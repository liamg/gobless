package main

import "github.com/liamg/gobless"

func main() {
	gui := gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	helloTextbox := gobless.NewTextBox()
	helloTextbox.Text = `Hello World!`
	helloTextbox.BorderColor = gobless.NewColor(0, 255, 0)
	helloTextbox.SetTitle("Message")

	quitTextbox := gobless.NewTextBox()
	quitTextbox.Text = `Press Ctrl-q to exit.`
	quitTextbox.Style = gobless.DefaultStyle
	quitTextbox.BorderColor = gobless.NewColor(255, 0, 0)

	otherTextbox := gobless.NewTextBox()
	otherTextbox.Text = `AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`

	row := gobless.NewRow(
		gobless.NewColumn(
			gobless.ColumnSizeOneSixth,
			helloTextbox,
		),
		gobless.NewColumn(
			gobless.ColumnSizeFiveSixths,
			quitTextbox,
		),
	)

	lowerRow := gobless.NewRow(
		gobless.NewColumn(
			gobless.ColumnSizeFull,
			otherTextbox,
		),
	)

	gui.Render(row, lowerRow)

	gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		gui.Close()
	})

	gui.Loop()
}
