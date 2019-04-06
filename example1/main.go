package main

import (
	"fmt"
)

var y = 1

func main() {
	x := 0
	increment := func() int {
		//fmt.Println(y)
		x++
		return x
	}
	fmt.Printf("%d LOL %t", increment(), increment())
	fmt.Println(x)
}
