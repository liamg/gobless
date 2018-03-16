package gobless

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/terminfo"
)

type GUI struct {
	width              int
	height             int
	keyEventChannel    chan KeyPressEvent
	resizeEventChannel chan ResizeEvent
	keyHandlers        map[Key][]func(event KeyPressEvent)
	resizeHandlers     []func(event ResizeEvent)
	renderMutex        sync.Mutex
	quitChan           chan bool
	closing            bool
	screen             tcell.Screen
	renderChannel      chan []Component
}

func NewGUI() *GUI {
	return &GUI{
		keyEventChannel:    make(chan KeyPressEvent, 1024),
		resizeEventChannel: make(chan ResizeEvent, 1024),
		keyHandlers:        map[Key][]func(event KeyPressEvent){},
		resizeHandlers:     []func(event ResizeEvent){},
		renderMutex:        sync.Mutex{},
		quitChan:           make(chan bool),
		renderChannel:      make(chan []Component, 1024),
	}
}

func (gui *GUI) Init() error {

	var err error
	gui.screen, err = tcell.NewScreen()
	if err != nil {
		if err == terminfo.ErrTermNotFound {
			return fmt.Errorf("Terminal entry not found. You can try setting:\n\n\texport TERM=xterm-256color\n\nOr simply:\n\n\texport TERM=xterm\n\n")
		}
		return err
	}
	if err := gui.screen.Init(); err != nil {
		return err
	}

	gui.screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	gui.screen.Clear()

	gui.width, gui.height = gui.screen.Size()

	return nil
}

func (gui *GUI) Size() (int, int) {
	return gui.width, gui.height
}

func (gui *GUI) Close() {
	if gui.closing {
		return
	}
	gui.closing = true
	close(gui.quitChan)
	gui.screen.Fini()
}

func (gui *GUI) Render(components ...Component) {
	gui.renderChannel <- components
}

func (gui *GUI) HandleKeyPress(key Key, handler func(event KeyPressEvent)) {
	_, ok := gui.keyHandlers[key]
	if !ok {
		gui.keyHandlers[key] = []func(event KeyPressEvent){}
	}

	gui.keyHandlers[key] = append(gui.keyHandlers[key], handler)
}

func (gui *GUI) HandleResize(handler func(event ResizeEvent)) {

	gui.resizeHandlers = append(gui.resizeHandlers, handler)
}

func (gui *GUI) Clear() {
	gui.screen.Clear()
}

func (gui *GUI) renderComponents(components []Component) {
	gui.renderMutex.Lock()
	defer gui.renderMutex.Unlock()

	rowCount := 0

	w, h := gui.Size()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gui.screen.SetCell(
				x,
				y,
				DefaultStyle.toTCell(),
				' ',
			)
		}
	}

	for _, component := range components {
		switch component.(type) {
		case *Row:
			rowCount++
		}
	}

	_, height := gui.Size()
	rowHeight := height
	spareHeight := 0
	if rowCount > 0 {
		rowHeightFloat := float64(height) / float64(rowCount)
		rowHeight = int(math.Floor(rowHeightFloat))
		spareHeight = height - (rowHeight * rowCount)
	}

	rowOffset := 0

	for _, component := range components {
		switch row := component.(type) {
		case *Row:
			row.width = w
			row.height = rowHeight + spareHeight // give "leftover" height runes to the first row, e.g. if there is a height of 10 runes and there are 3 rows, each row will get 3 runes of height, and the first row will get the "spare" 1, to make 10.
			row.y = rowOffset
			row.x = 0
			rowOffset += row.height
			spareHeight = 0
		}
		for _, tile := range component.GetTiles(gui) {
			for point, cell := range tile.Cells {
				if point.X >= 0 && point.X <= tile.Rectangle.Max.X-tile.Rectangle.Min.X && point.Y >= 0 && point.Y <= tile.Rectangle.Max.Y-tile.Rectangle.Min.Y {
					gui.screen.SetCell(
						tile.Rectangle.Min.X+point.X,
						tile.Rectangle.Min.Y+point.Y,
						cell.Style.toTCell(),
						cell.Rune,
					)
				}
			}
		}

	}

	gui.screen.Show()

	gui.width, gui.height = gui.screen.Size()
}
func (gui *GUI) Loop() {

	tcellEventChan := make(chan tcell.Event)

	// tcell event polling on a seperate routine, because it blocks and you can't pass in a context :/
	go func() {
		for {
			tcellEventChan <- gui.screen.PollEvent()
		}
	}()

	for {
		select {
		case <-gui.quitChan:
			time.Sleep(time.Millisecond * 100) // fix weird RC where terminal is orrupt unless we let tcell tidy up a tiny bit longer :/
			return

		case keyEvent := <-gui.keyEventChannel:
			handlers, ok := gui.keyHandlers[Key(keyEvent.Key)]
			if ok {
				for _, handler := range handlers {
					go handler(keyEvent)
				}
			}
		case resizeEvent := <-gui.resizeEventChannel:

			gui.width, gui.height = gui.screen.Size()

			for _, handler := range gui.resizeHandlers {
				go handler(resizeEvent)
			}

		case ev := <-tcellEventChan:
			switch event := ev.(type) {
			case *tcell.EventKey:
				gui.keyEventChannel <- KeyPressEvent{
					Key: Key(event.Key()),
				}
			case *tcell.EventResize:
				w, h := event.Size()
				gui.resizeEventChannel <- ResizeEvent{
					Width:  w,
					Height: h,
				}
			default:
				//fmt.Printf("Event: %#v\n", event)
			}

		case components := <-gui.renderChannel:
			gui.renderComponents(components)
		}
	}

}
