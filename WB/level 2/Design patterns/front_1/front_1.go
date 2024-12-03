package main 

import (
	"fmt"
)

type Device interface{
	TurnOff()
	TurnOn()
}

type Light struct{
	status bool
}

func (l *Light) TurnOff() {
	l.status = false
	fmt.Println("Light is turn off")
}

func (l *Light) TurnOn() {
	l.status = true
	fmt.Println("Light is turn on")
}

type Heating struct{
	status bool
}

func (h *Heating) TurnOff() {
	h.status = false
	fmt.Println("Heating is turn off")
}

func (h *Heating) TurnOn() {
	h.status = true
	fmt.Println("Heating is turn on")
}

type SecuritySystem struct{
	status bool
}

func (s *SecuritySystem) TurnOff() {
	s.status = false
	fmt.Println("Security system is turn off")
}

func (s *SecuritySystem) TurnOn() {
	s.status = true
	fmt.Println("Security system is turn on")
}

type SmartHomeFacade struct{
	light Device
	heating Device
	securitySystem Device
}

func NewSmartHomeFacade() *SmartHomeFacade{
	return &SmartHomeFacade{
		light: &Light{},
		heating: &Heating{},
		securitySystem: &SecuritySystem{},
	}
}

func (s *SmartHomeFacade) LeaveHome() {
	s.light.TurnOff()
	s.heating.TurnOff()
	s.securitySystem.TurnOn()
}

func (s *SmartHomeFacade) ArriveHome(){
	s.light.TurnOn()
	s.heating.TurnOn()
	s.securitySystem.TurnOff()
}




func main(){
	status := NewSmartHomeFacade()

	status.LeaveHome()

	status.ArriveHome()
}