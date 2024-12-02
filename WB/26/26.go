package main

import (
	"fmt"
)

func main(){
	str := "aabcd"
	var count int
	for _, i := range str{
		for _, j := range str{
			if i == j{
				count++
			}

		}
		if count == 2 {
			fmt.Print("false")
			return
		}else{
			count = 0
		}
	}
	fmt.Print("true")
}