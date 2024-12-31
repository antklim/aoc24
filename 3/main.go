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

	operands, err := ReadMul(f)
	if err != nil {
		fmt.Println("failed to read input file:", err)
		return
	}

	operations, _ := ParseOperands(operands)

	result := 0
	for _, o := range operations {
		result += o[0] * o[1]
	}

	fmt.Println("multiplication result:", result)
}
