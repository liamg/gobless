package gobless

import (
	"math"
	"sync"
	"time"

	"github.com/gdamore/tcell"
)

type GUI struct {
	width           int
	height          int
	keyEventChannel chan KeyPressEvent
	keyHandlers     map[Key][]func(event KeyPressEvent)
	renderMutex     sync.Mutex
	quitChan        chan bool
	closing         bool
	screen          tcell.Screen
	renderChannel   chan []Component
}

func NewGUI() *GUI {
	return &GUI{
		keyEventChannel: make(chan KeyPressEvent, 1024),
		keyHandlers:     map[Key][]func(event KeyPressEvent){},
		renderMutex:     sync.Mutex{},
		quitChan:        make(chan bool),
		renderChannel:   make(chan []Component, 1024),
	}
}

func (gui *GUI) Init() error {

	var err error
	gui.screen, err = tcell.NewScreen()
	if err != nil {
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

func (gui *GUI) Clear() {
	gui.screen.Clear()
}

func (gui *GUI) renderComponents(components []Component) {
	gui.renderMutex.Lock()
	defer gui.renderMutex.Unlock()

	rowCount := 0

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
			row.height = rowHeight + spareHeight // give "leftover" height runes to the first row, e.g. if there is a height of 10 runes and there are 3 rows, each row will get 3 runes of height, and the first row will get the "spare" 1, to make 10.
			row.y = rowOffset
			row.x = 0
			rowOffset += row.height
			spareHeight = 0
		}
		for _, tile := range component.GetTiles(gui) {
			for point, cell := range tile.Cells {
				gui.screen.SetCell(
					tile.Rectangle.Min.X+point.X,
					tile.Rectangle.Min.Y+point.Y,
					cell.Style.toTCell(),
					cell.Rune,
				)
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

		case ev := <-tcellEventChan:
			switch event := ev.(type) {
			case *tcell.EventKey:
				gui.keyEventChannel <- KeyPressEvent{
					Key: Key(event.Key()),
				}
			}

		case keyEvent := <-gui.keyEventChannel:
			handlers, ok := gui.keyHandlers[Key(keyEvent.Key)]
			if ok {
				for _, handler := range handlers {
					go handler(keyEvent)
				}
			}
		case components := <-gui.renderChannel:
			gui.renderComponents(components)
		}
	}

}
