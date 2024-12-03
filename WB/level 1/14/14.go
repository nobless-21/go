package main

import(
	"fmt"
	"reflect"
)

func main(){
	var a interface{}
	ch := make(chan int)
	a = ch
	type_func(a)
}

func type_func(a interface{}){
	switch reflect.TypeOf(a).Kind(){
		case reflect.Int:
			fmt.Println("переменна типа int")
		case reflect.String:
			fmt.Println("переменная типа string")
		case reflect.Bool:
			fmt.Println("переменная типа bool")
		case reflect.Chan:
			fmt.Println("переменная типа канал")
	}
}