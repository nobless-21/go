package main

import (
	"fmt"
)

func main() {
	arr := [5]string{"cat", "cat", "dog", "cat", "tree"}
	mapa := make(map[string]bool)
	for _, i := range arr {
		mapa[i] = true
	}
	fmt.Println(mapa)
}