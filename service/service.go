package service

import (
	"fmt"
	"log"
	"time"
)

type commonFn func() error

func DoService(validate commonFn, inputMapping commonFn, business commonFn) {
	defer func(s time.Time) {
		log.Printf("elpased time %0.2d ns", time.Since(s).Nanoseconds())
	}(time.Now())
	validate()
	inputMapping()
	business()
}

type Service interface {
	Echo() string
	//Speak() string
}

type DemoService struct {
}

func (s *DemoService) Validate() error {
	fmt.Println("Validate")
	return nil
}

func (s *DemoService) InputMapping() error {
	fmt.Println("InputMapping")
	return nil
}

func (s *DemoService) Business() error {
	fmt.Println("Business")
	return nil
}

func (s *DemoService) Echo() string {
	DoService(s.Validate, s.InputMapping, s.Business)
	return "ECHOOOOOO"
}
