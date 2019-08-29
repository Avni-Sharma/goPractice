package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("newfile")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	data, _ := reader.Peek(5) // reader data is a byte arr
	fmt.Println(string(data))
	file.Close()
}
