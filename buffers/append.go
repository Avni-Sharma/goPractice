package main

import (
	"os"
)

func main() {
	file, err := os.OpenFile("newfile", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	file.WriteString("I have been appended\n")
	file.Close()
}
