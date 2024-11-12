package main

import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	var mu sync.Mutex

	var arr = [5]int{2,4,6,8,10} 
	sum := 0

	result := make(chan int, len(arr))

	for _, i := range arr {
		wg.Add(1)
		go func(num int){
			defer wg.Done()
			out := sqr(num)
			result <- out
		}(i)
	}

	go func(){
		wg.Wait()
		close(result)
	}()


	for result := range result {
		mu.Lock()
		sum += result
		mu.Unlock()
	}

	fmt.Println(sum)
}

func sqr(a int) int{
	return a * a
}