package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	path := dir + "/" + os.Getenv("GOFILE")
	fmt.Println("path", path)
	fmt.Println("line", os.Getenv("GOLINE"))
	fmt.Println("package", os.Getenv("GOPACKAGE"))

	f, _ := os.Open(path)
	defer f.Close()

	// data, _ := io.ReadAll(f)
	// fmt.Println(string(data))

	scanner := bufio.NewReaderSize(f, 128)
	for {
		buf, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(buf))
	}
}
