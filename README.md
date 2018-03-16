# gobless
Build beautiful terminal dashboards and GUIs in Golang. 

*Requires golang >=1.10*

The name was intended to be pronounced _go-bless_ due to the inspiration from [blessed-contrib](https://github.com/yaronn/blessed-contrib), but _gob-less_ is definitely funnier, so we'll go with that.

_screenshot here plz_

## Get Started

You can get started by viewing the various [examples](_examples/).

## Components

### Text Box

### Bar Chart

### Sparkline

### Progress Bar

### Sparklines

### Line Chart

### Block Chart

(kind of like blessed-contrib's stacked gauge)

### Donut

### Log

### Table

### Tree

## Layout System

Gobless includes a built-in CSS style 12 column nestable grid layout system. This uses `Row`s and `Column`s (which are themselves components) to quickly build a UI with minimal effort. 

```golang
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
```

You can also detect [resize events](#events) and automatically resize the layout of your components if you wish.

```golang
gui.HandleResize(func(event gobless.ResizeEvent) {
	gui.Render(rows...)
})
```

A [full example](_examples/gridlayout) is also available.

Alternatively, if you need to, you can position components absolutely using the `.SetX()`, `.SetY()`, `.SetWidth()` and `.SetHeight()` methods, omitting the presence of rows/columns entirely.

## Events

Gobless can currently detect two types of events:

### Key Press Events

You can add your own handlers for key presses, utilising the `gobless.Key*` constants.

For example, to stop rendering the GUI when the user presses `CTRL` + `Q`, you could do:

```
gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent){
		gui.Close()
})
```

## Terminology

- A *Cell* refers to an individual terminal cell, in which a single character can be drawn.
- A *Tile* refers to a rectangular grouping of cells.
- A *Component* is an entity which provides a tile - or tiles - to the GUI to render. A bar chart, for example.
