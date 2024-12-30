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
	//	operands, err := ParseInput(f)
	if err != nil {
		fmt.Println("failed to read input file:", err)
		return
	}

	operations, err := ParseOperands(operands)

	result := 0
	for _, o := range operations {
		fmt.Println(o[0], o[1])
		result += o[0] * o[1]
	}

	fmt.Println("multiplication result:", result)
	//fmt.Println("errors:", err)
}
