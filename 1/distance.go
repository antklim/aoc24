package main

import (
	"errors"
	"sort"
)

var ErrDifferentLength = errors.New("lists have different lengths")

func Distance(a, b []int) (int, error) {
	if len(a) != len(b) {
		return 0, ErrDifferentLength
	}

	sort.Ints(a)
	sort.Ints(b)

	result := 0
	for i := 0; i < len(a); i++ {
		x := 0
		if a[i] > b[i] {
			x = a[i] - b[i]
		} else {
			x = b[i] - a[i]
		}

		result += x
	}

	return result, nil
}
