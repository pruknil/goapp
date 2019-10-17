package service

type Service interface {
	Echo() string
}

type DemoService struct {
}

func (s *DemoService) Echo() string {
	//fmt.Println("Echo")
	return "ECHOOOOOO"
}
