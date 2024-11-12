package main

import (
	"fmt"
	"unsafe"
)

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func main() {
	a := int64(1)
	fmt.Println("memory pointer for var `a`", unsafe.Pointer(&a))
	fmt.Println("memory size for var `a`", unsafe.Sizeof(a))

	println("-------")

	f := 3.0
	fmt.Printf("%d\n", Float64bits(f))
	fmt.Printf("%#016x\n", Float64bits(f))
	fmt.Printf("%b\n", Float64bits(f))

	// return
	println("-------")

	type Message struct {
		name  string // 8 b
		flag1 bool   //1 b
		flag2 bool   // 1 b
	}

	msg := Message{
		flag1: false,
		name:  "sadadsasdasdasdadsasdadsooioiodasdadasd",
		flag2: false,
	}

	fmt.Println("memory size for Message struct", unsafe.Sizeof(msg))

	fmt.Println(
		"flag1 Sizeof", unsafe.Sizeof(msg.flag1),
		"Alignof", unsafe.Alignof(msg.flag1),
		"Offsetof", unsafe.Offsetof(msg.flag1),
	)

	fmt.Println(
		"flag2 Sizeof", unsafe.Sizeof(msg.flag2),
		"Alignof", unsafe.Alignof(msg.flag2),
		"Offsetof", unsafe.Offsetof(msg.flag2),
	)

	fmt.Println(
		"name Sizeof", unsafe.Sizeof(msg.name),
		"Alignof", unsafe.Alignof(msg.name),
		"Offsetof", unsafe.Offsetof(msg.name),
	)

}
