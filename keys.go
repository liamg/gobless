package gobless

import (
	"github.com/gdamore/tcell"
)

type Key uint16

// Abstract from tcell
const (
	KeyF1     = Key(tcell.KeyF1)
	KeyF2     = Key(tcell.KeyF2)
	KeyF3     = Key(tcell.KeyF3)
	KeyF4     = Key(tcell.KeyF4)
	KeyF5     = Key(tcell.KeyF5)
	KeyF6     = Key(tcell.KeyF6)
	KeyF7     = Key(tcell.KeyF7)
	KeyF8     = Key(tcell.KeyF8)
	KeyF9     = Key(tcell.KeyF9)
	KeyF10    = Key(tcell.KeyF10)
	KeyF11    = Key(tcell.KeyF11)
	KeyF12    = Key(tcell.KeyF12)
	KeyInsert = Key(tcell.KeyInsert)
	KeyDelete = Key(tcell.KeyDelete)
	KeyHome   = Key(tcell.KeyHome)
	KeyEnd    = Key(tcell.KeyEnd)
	KeyPgUp   = Key(tcell.KeyPgUp)
	KeyPgDn   = Key(tcell.KeyPgDn)
	KeyUp     = Key(tcell.KeyUp)
	KeyDown   = Key(tcell.KeyDown)
	KeyLeft   = Key(tcell.KeyLeft)
	KeyRight  = Key(tcell.KeyRight)
)

const (
	KeyCtrlSpace      = Key(tcell.KeyCtrlSpace)
	KeyCtrlA          = Key(tcell.KeyCtrlA)
	KeyCtrlB          = Key(tcell.KeyCtrlB)
	KeyCtrlC          = Key(tcell.KeyCtrlC)
	KeyCtrlD          = Key(tcell.KeyCtrlD)
	KeyCtrlE          = Key(tcell.KeyCtrlE)
	KeyCtrlF          = Key(tcell.KeyCtrlF)
	KeyCtrlG          = Key(tcell.KeyCtrlG)
	KeyBackspace      = Key(tcell.KeyBackspace)
	KeyCtrlH          = Key(tcell.KeyCtrlH)
	KeyTab            = Key(tcell.KeyTab)
	KeyCtrlI          = Key(tcell.KeyCtrlI)
	KeyCtrlJ          = Key(tcell.KeyCtrlJ)
	KeyCtrlK          = Key(tcell.KeyCtrlK)
	KeyCtrlL          = Key(tcell.KeyCtrlL)
	KeyEnter          = Key(tcell.KeyEnter)
	KeyCtrlM          = Key(tcell.KeyCtrlM)
	KeyCtrlN          = Key(tcell.KeyCtrlN)
	KeyCtrlO          = Key(tcell.KeyCtrlO)
	KeyCtrlP          = Key(tcell.KeyCtrlP)
	KeyCtrlQ          = Key(tcell.KeyCtrlQ)
	KeyCtrlR          = Key(tcell.KeyCtrlR)
	KeyCtrlS          = Key(tcell.KeyCtrlS)
	KeyCtrlT          = Key(tcell.KeyCtrlT)
	KeyCtrlU          = Key(tcell.KeyCtrlU)
	KeyCtrlV          = Key(tcell.KeyCtrlV)
	KeyCtrlW          = Key(tcell.KeyCtrlW)
	KeyCtrlX          = Key(tcell.KeyCtrlX)
	KeyCtrlY          = Key(tcell.KeyCtrlY)
	KeyCtrlZ          = Key(tcell.KeyCtrlZ)
	KeyEsc            = Key(tcell.KeyEsc)
	KeyCtrlBackslash  = Key(tcell.KeyCtrlBackslash)
	KeyCtrlUnderscore = Key(tcell.KeyCtrlUnderscore)
	KeyBackspace2     = Key(tcell.KeyBackspace2)
)
