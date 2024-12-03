package main

import (
	"fmt"
)

func main(){
	var a int
	fmt.Scan(&a)
	slice := make([]int, 20)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	var newSlice []int
	newSlice = append(newSlice, slice[:a]...)
	newSlice = append(newSlice, slice[a+1:]...)
	fmt.Print(newSlice)
}