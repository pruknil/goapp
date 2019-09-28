package backends

type Backener interface {
	Open()
	Close()
}
