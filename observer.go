package main

type subject interface {
	registerObserver(o observer)
	removeObserver(o observer)
	notifyObservers()
}

type observer interface {
}

type httpRouter struct {
}

func (h httpRouter) registerObserver(o observer) {
	panic("implement me")
}

func (h httpRouter) removeObserver(o observer) {
	panic("implement me")
}

func (h httpRouter) notifyObservers() {
	panic("implement me")
}
