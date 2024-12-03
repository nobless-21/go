package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		work()
	}()

	time.Sleep(time.Second * 5)
	wg.Wait()
	fmt.Println("main закончил работу")
}

func work(){
	for i:= 0; i<5;i++{
		fmt.Println("горутина работает")
		time.Sleep(time.Second)
	}
	fmt.Println("горутина завершила работу")
}