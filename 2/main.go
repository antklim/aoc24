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

	safeReportsAmount := 0
	var unsafeReports []Report
	for _, report := range reports {
		if report.IsSafe() {
			safeReportsAmount++
		} else {
			unsafeReports = append(unsafeReports, report)
		}
	}

	fmt.Println("safe reports amount:", safeReportsAmount)

	for _, report := range unsafeReports {
		if report.IsSafeWithDampener() {
			safeReportsAmount++
		}
	}

	fmt.Println("safe reports amount with dampener:", safeReportsAmount)
}
