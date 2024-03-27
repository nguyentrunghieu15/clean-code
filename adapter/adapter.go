package main

import "fmt"

type BirdInterface interface {
	Fly()
}

type Bird struct {
	Name string
}

func (s *Bird) Fly() {
	fmt.Printf("A %s is flying\n", s.Name)
}

type TerrestrialAnimalInterface interface {
	Run()
}

type TerrestrialAnimal struct {
	Name string
}

func (s *TerrestrialAnimal) Run() {
	fmt.Printf("A %s is running\n", s.Name)
}

type BirdAdapter struct {
	Bird
}

func (s *BirdAdapter) Run() {
	fmt.Printf("A %s is running\n", s.Name)
}

func main() {
	wild_duck := BirdAdapter{Bird: Bird{Name: " Wild Duck"}}
	wild_duck.Run()
	wild_duck.Fly()
}
