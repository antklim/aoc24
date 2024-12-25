package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseInput(r io.Reader) ([]Report, error) {
	var reports []Report

	i := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		i++
		s := strings.Split(scanner.Text(), " ")

		var a []int
		for j, v := range s {
			if x, err := strconv.Atoi(strings.TrimSpace(v)); err != nil {
				return nil, fmt.Errorf("failed to parse %d item on line %d: %w", j, i, err)
			} else {
				a = append(a, x)
			}
		}

		reports = append(reports, NewReport(a))
	}

	return reports, scanner.Err()
}
