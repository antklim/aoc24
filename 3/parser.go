package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseOperands(operands []string) ([][]int, error) {
	var errs error
	var result [][]int

	for _, o := range operands {
		v := strings.Split(o, ",")
		if len(v) != 2 {
			errs = errors.Join(errs, fmt.Errorf("invalid operand %q", v))
		} else {
			a, b := v[0], v[1]

			if len(a) < 1 || len(a) > 3 {
				errs = errors.Join(errs, fmt.Errorf("invalid operand %q: first argument either too long or short: %q", v, a))
			}
			if len(b) < 1 || len(b) > 3 {
				errs = errors.Join(errs, fmt.Errorf("invalid operand %q: second argument either too long or short: %q", v, b))
			}

			if len(a) >= 1 && len(a) <= 3 && len(b) >= 1 && len(b) <= 3 {
				ai, aerr := strconv.Atoi(a)
				bi, berr := strconv.Atoi(b)

				if aerr != nil {
					errs = errors.Join(errs, fmt.Errorf("invalid operand %q: invalid first argument: %w", v, aerr))
				}
				if berr != nil {
					errs = errors.Join(errs, fmt.Errorf("invalid operand %q: invalid second argument: %w", v, berr))
				}
				if aerr == nil && berr == nil {
					result = append(result, []int{ai, bi})
				}
			}
		}
	}

	return result, errs
}
