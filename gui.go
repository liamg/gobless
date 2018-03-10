package gobless

import (
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

	for _, component := range components {
		for _, tile := range component.GetTiles() {
			for point, cell := range tile.Cells {
				if point.Add(tile.Rectangle.Min).In(tile.Rectangle) {
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
