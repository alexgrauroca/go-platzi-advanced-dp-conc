package observer

type Observer interface {
	getId() string
	updateValue(string)
}
