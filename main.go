package main

import (
	"fmt"
)

func hello() func() string {

	return func() string {
		fmt.Println("World")
		return "hey"
	}
}
func main() {
	arre_re := hello()
	fmt.Println(arre_re())

}
