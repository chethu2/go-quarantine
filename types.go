package main

import (
	"fmt"
	"math/cmplx"
)

var (
	a      int        = 10
	flag   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T, Value: %v\n", a, a)
	fmt.Printf("Type: %T, Value: %v\n", flag, flag)
	fmt.Printf("Type: %T, Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)

}
