package main 

import (
	"fmt"
	"strings"
)

func main(){
	var str strings.Builder

	str.WriteString("hello")
	str.WriteString(" ")
	str.WriteString("world")

	result := str.String()

	fmt.Println(result)
}