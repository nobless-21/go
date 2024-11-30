package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	words := strings.Fields(str)
	for i := 0; i < len(words)/ 2; i++ {
		words[i], words[len(words)- 1 - i] = words[len(words)- 1 - i], words[i]
	}
	fmt.Println(strings.Join(words, " "))
}

