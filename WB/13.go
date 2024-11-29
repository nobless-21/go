package main

import (
	"fmt"
)

func main() {
	a,b := 3,4
	fmt.Println(a,b)
	swap(&a,&b)
	fmt.Println(a,b)
}

func swap(a,b *int){
	*a,*b = *b,*a
}