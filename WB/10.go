package main

import (
	"fmt"
)

func main() {
	arr := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var slise1 []float32
	var slise2 []float32
	var slise3 []float32
	var slise4 []float32
	for _, i:= range arr {
		if i > 30 {
			slise4 = append(slise4, i)
		} else if i > 20 && i < 30 {
			slise3 = append(slise3, i)
		} else if i > 10 && i < 20 {
			slise2 = append(slise2, i)
		} else if i > -30 && i < -20 {
			slise1 = append(slise1, i)
		}
	}
	fmt.Println("-20", slise1)
	fmt.Println("10", slise2)
	fmt.Println("20", slise3)
	fmt.Println("30", slise4)
}