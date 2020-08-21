package main

import "fmt"

const p string = "I am a string constant"
const i = 7

func main() {
	fmt.Println("hello to constants")
	fmt.Println("%s", p)
	fmt.Println("%i", i)
	fmt.Println()
	fmt.Printf("%s", p)
	fmt.Println()
	fmt.Printf("%i", i)
}
