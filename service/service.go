package service

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type commonFn func() error

func DoService(service ServiceTemplate) error {
	defer func(s time.Time) {
		log.Printf("elpased time %0.2d ns", time.Since(s).Nanoseconds())
	}(time.Now())

	if service.Validate() != nil {
		errors.New("Validate Error")
	}

	service.InputMapping()
	service.Business()
	service.OutputMapping()

	return nil
}

type Service interface {
	Echo() string
	//Speak() string
}

type ServiceTemplate interface {
	Validate() error
	OutputMapping() error
	InputMapping() error
	Business() error
}

type DemoService struct {
	ServiceTemplate
}

func (s *DemoService) Validate() error {
	fmt.Println("Validate")
	return nil
}

func (s *DemoService) OutputMapping() error {
	fmt.Println("OutputMapping")
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
	DoService(s)
	return "ECHOOOOOO"
}
