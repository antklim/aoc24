package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseInput(r io.Reader) ( /* a */ []int /* b */, []int /* err */, error) {
	var a, b []int

	i := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		i++
		s := strings.SplitN(scanner.Text(), " ", 2)

		if x, err := strconv.Atoi(strings.TrimSpace(s[0])); err != nil {
			return nil, nil, fmt.Errorf("failed to parse first argument on line %d: %w", i, err)
		} else {
			a = append(a, x)
		}

		if x, err := strconv.Atoi(strings.TrimSpace(s[1])); err != nil {
			return nil, nil, fmt.Errorf("failed to parse second argument on line %d: %w", i, err)
		} else {
			b = append(b, x)
		}
	}

	return a, b, scanner.Err()
}
