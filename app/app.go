package app

type Config struct {
	Backend
}
type Backend struct {
	Hsm
}
type Hsm struct {
	Host string
	Port string
}
