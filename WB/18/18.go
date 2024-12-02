package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct{
	score int
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count counter
	for i := 0;i < 10;i++ {
		wg.Add(1)
		go work(&count, &wg, &mu)
		time.Sleep(time.Second)
	}

	wg.Wait()
	fmt.Print((count.score))
}

func work(count *counter, wg *sync.WaitGroup, mu *sync.Mutex){
	defer wg.Done()
	mu.Lock()
	count.score++
	fmt.Println(count.score)
	mu.Unlock()
}