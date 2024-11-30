package main
 import (
	"fmt"
	"sort"
 )

 func main() {
	target := 16
	slice := []int{3, 7, 12, 5, 18, 1, 14, 9, 6, 11, 2, 15, 4, 8, 17, 10, 13, 16, 2, 7, 5, 14, 11, 3, 9}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	index := sort.Search(len(slice), func(i int)bool {
		return slice[i] >= target
	} )
	
	if slice[index] == target {
		fmt.Printf("индекс - %d\n число - %d", index, slice[index])
	}
 }