package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("newfile")
	if err != nil {
		panic(err)
	}
	buf := bufio.NewWriter(file)                // Create a buffer for the OS File
	buf.WriteString("Writing data to buffer\n") // Write data to buffer
	buf.Flush()                                 // Clears the buffer data and writes to the file
	file.Close()
}
