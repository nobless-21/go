package main

import (
    "fmt"
)

func main(){
    arr1 := [10]int{1,2,3,4,5,6,7,8,9,10}
    arr2 := [10]int{6,7,8,9,10,11,2,13,14,15}

    work(arr1,arr2)
}

func work(arr1, arr2 [10]int) {
    var slise []int
    var score int
    for i := 0;i < len(arr1); i++ {
        score = 0
        for t := 0; t < len(arr2); t++ {
            if arr1[i] == arr2[t] && score == 0 {
                slise = append(slise, arr1[i])
                score = 1
            }
        }
    }
    fmt.Println(slise)
}