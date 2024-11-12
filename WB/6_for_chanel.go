package main

import (
	"fmt"
	"time"
)

func main(){
	stopChan := make(chan struct{})

	go work(stopChan)

	time.Sleep(5*time.Second)

	close(stopChan)
	time.Sleep(time.Second)
	fmt.Println("main завершился")
}

func work(stopChan <-chan struct{}){
	for {
		select{
		case <-stopChan:
			fmt.Println("горутина остановлена")
			return
		default:
			fmt.Println("горутина работает")
			time.Sleep(time.Second)
		}
	}
}