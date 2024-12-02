package main

import (
	"fmt"
	"math/big"
)

func main(){
	bigInt1 := big.NewInt(1 << 22)
	bigInt2 := big.NewInt(1 << 22)

	sum := new(big.Int).Add(bigInt1, bigInt2)
	sub := new(big.Int).Sub(bigInt1, bigInt2)
	mul := new(big.Int).Mul(bigInt1, bigInt2)
	div := new(big.Int).Div(bigInt1, bigInt2)

	fmt.Println(sum, sub, mul, div)
}