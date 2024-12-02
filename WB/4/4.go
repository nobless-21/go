package main 

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var countWorkers int
	maxBuffer := 10000
	fmt.Scan(&countWorkers)

	input := make(chan int, maxBuffer)
	result := make(chan int, maxBuffer)
	
	for i:= 0; i< countWorkers; i++ {
		wg.Add(1)
		go worker(input, result, &wg)
	}

	for i := 0;;i++ {
		input<- i
		if i+1 == maxBuffer {
			break
		}
	}
	close(input)

	go func(){
		wg.Wait()
		close(result)
	}()

	for res := range result{
		fmt.Println(res)
	}
}

func worker (input <-chan int, result chan<- int, wg *sync.WaitGroup){
	defer wg.Done()
	for in := range input{
		time.Sleep(time.Second)
		result <- in
	}
}