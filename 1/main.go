package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("failed to open input file:", err)
		return
	}

	a, b, err := ParseInput(f)
	if err != nil {
		fmt.Println("failed to parse input file:", err)
		return
	}

	d, err := Distance(a, b)
	if err != nil {
		fmt.Println("failed to calculate distance:", err)
		return
	}

	fmt.Println("distance between lists:", d)
}
