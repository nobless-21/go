package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	mapa := make(map[int]int)
	var arr [10]int

	for i:= 0; i< len(arr); i++ {
		arr[i] = i
	}

	for i:= range arr {
		wg.Add(1)
		go work(&mapa, i, &wg, &mu)
	}
	wg.Wait()
	for key, value:= range mapa {
		fmt.Println(key , value)
	}
}

func work(mapa *map[int]int, a int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	(*mapa)[a] = a*a
}