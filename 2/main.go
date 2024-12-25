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

	reports, err := ParseInput(f)
	if err != nil {
		fmt.Println("failed to parse input file:", err)
		return
	}

	safeReports := 0
	for _, report := range reports {
		if report.IsSafe() {
			safeReports++
		}
	}

	fmt.Println("safe reports amount:", safeReports)
}
