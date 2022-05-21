package observer

type Topic interface {
	register(observer Observer)
	broadcast()
}
