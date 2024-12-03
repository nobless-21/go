package main

import(
	"fmt"
	"time"
)

func main(){
	fmt.Println("start")
	sleep(time.Second * 3)
	fmt.Println("end")
}

func sleep(d time.Duration){
	done := time.After(d)
	<-done
}