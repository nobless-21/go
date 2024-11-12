package main

import (
	"fmt"
	"sync"
)

func main() {
	var a = [5]int{2,4,6,8,10}
	var wg sync.WaitGroup
	for _, i := range a {
		wg.Add(1)
		go sqr(i, &wg)
	}
	wg.Wait()
}

func sqr(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(a * a)
}