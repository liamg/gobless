package gobless

type Event interface{}

type KeyPressEvent struct {
	Key Key
	//Modifier Modifier // @todo
}
