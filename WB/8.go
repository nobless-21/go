package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a,b int64
	fmt.Println("введите число")
	fmt.Scan(&a)
	fmt.Println("введите номер бита, который хотите поменять")
	fmt.Scan(&b)
	s := strconv.FormatInt(a, 2)
	fmt.Println(s)
	a ^= (1 << b)
	s1 := strconv.FormatInt(a,2)
	fmt.Println(s1)
}