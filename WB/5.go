package main

import (
	"fmt"
)

func main(){
	arr := [5]int{1,2,3,4,5}
	input := make(chan int, len(arr))
	for _, i := range arr{
		input <- i
	}
	close(input)
	channel(input)
}

func channel (input <-chan int){
	for i := range input {
		fmt.Println(i)
	}
}