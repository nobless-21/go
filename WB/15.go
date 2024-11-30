package main

import (
	"fmt"
)

func main() {
	justString := someFunc()
	fmt.Print(justString)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func createHugeString(size int) string {
	return string(make([]byte, size))
}
