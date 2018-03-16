package gobless

type Event interface{}

type KeyPressEvent struct {
	Key Key
	//Modifier Modifier // @todo
}

type ResizeEvent struct {
	Width  int
	Height int
}
