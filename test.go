package main

import (
	"fmt"
	"unsafe"
)

type S2 struct { // align 16
	a string // size 16 = 8 + 8
	d int    // size 8
	c string // size 16 = 8 + 8
}

func main() {
	y := S2{}
	fmt.Println(unsafe.Sizeof(y))

	fmt.Println(unsafe.Offsetof(y.a))
	fmt.Println(unsafe.Offsetof(y.d))
	fmt.Println(unsafe.Offsetof(y.c))

	fmt.Println(&y.a)
	fmt.Println(&y.d)
	fmt.Println(&y.c)
}
