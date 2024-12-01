package main

import (
	"fmt"
)

type Movable interface{
	Move() string
}

type Bird struct{}

type BirdAdapter struct{
	bird *Bird
}

type Fish struct{}

type FishAdapter struct{
	fish *Fish
}

type Dog struct{}

type DogAdapter struct{
	dog *Dog
}

func (b *Bird) fly() string{
	return "the bird flying"
}

func (b *BirdAdapter) Move() string{
	return b.bird.fly()
}

func (f *Fish) swim() string{
	return "the fich is swimming"
}

func (f *FishAdapter) Move() string {
	return f.fish.swim()
}

func (d *Dog) run() string{
	return "the dog is running"
}

func (d *DogAdapter) Move() string{
	return d.dog.run()
}



func main(){
	bird := &Bird{}
	fish := &Fish{}
	dog := &Dog{}

	birdAdapter := &BirdAdapter{bird: bird}
	fishAdapter := &FishAdapter{fish: fish}
	dogAdapter := &DogAdapter{dog: dog}

	animals := []Movable{birdAdapter, fishAdapter,dogAdapter}
	for _, animal := range animals {
		fmt.Println(animal.Move())
	}
}