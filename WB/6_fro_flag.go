package main

import (
	"fmt"
	"time"
)

func main() {
	var flag bool

	go work(&flag)

	time.Sleep(5*time.Second)
	flag = true
	time.Sleep(time.Second)
	fmt.Println("main завершил работу")
}

func work(flag *bool) {
	for {
		if *flag == true {
			fmt.Println("горутина завершила свою работу")
			return
		}else{
			fmt.Println("горутина работает")
			time.Sleep(time.Second)
		}
	}
}