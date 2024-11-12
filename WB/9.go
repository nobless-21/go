package main

import (
	"fmt"
	"sync"
	"math"
)

func main() {
	var wg sync.WaitGroup

	var arr [10]int
	for i := 0; i < len(arr); i++{
		arr[i] = i
	}

	chanel1 := make(chan int)
	chanel2 := make(chan float64)

	go func() {
		defer close(chanel1)
		defer wg.Wait()
		for i := range arr {
			wg.Add(1)
			go work1(i, chanel1, &wg)
		}
	}()

	go func() {
		defer close(chanel2)
		defer wg.Wait()
		for a := range chanel1 {
			wg.Add(1)
			go work2(a, chanel2, &wg)
		}
	}()

	for val1 := range chanel2 {
		fmt.Println(val1)
	}
}

func work1(i int, chanel1 chan<- int, wg *sync.WaitGroup){
	defer wg.Done()
	chanel1 <- i
}

func work2(a int, chanel2 chan<- float64, wg *sync.WaitGroup){
	defer wg.Done()
	i := a
	chanel2 <- math.Pow(float64(i),2)
}