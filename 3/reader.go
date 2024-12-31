package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Operation int

const (
	OperationUnknown Operation = iota - 1
	OperationDo
	OperationDont
	OperationMul
)

var operations = []string{"do()", "don't()", "mul("}

// NextOperation reads from the reader until the next operation found or reached the end of the reader.
func NextOperation(r *bufio.Reader) (Operation, error) {
	v := ""

	c, _, err := r.ReadRune()
	for err == nil {
		v += string(c)

		idx := findIdx(operations, func(s string) bool {
			return v == s
		})
		if idx >= 0 {
			return Operation(idx), nil
		}

		idx = findIdx(operations, func(s string) bool {
			return strings.HasPrefix(s, v)
		})
		if idx == -1 {
			v = ""
		}

		c, _, err = r.ReadRune()
	}

	return OperationUnknown, err
}

func ReadMul(r io.Reader) ([]string, error) {
	reader := bufio.NewReader(r)

	var result []string
	do := true

	o, err := NextOperation(reader)
	for err == nil {
		switch o {
		case OperationDo:
			do = true
		case OperationDont:
			do = false
		case OperationMul:
			if !do {
				break
			}

			operands, valid, done, err := readOperands(reader)
			if err != nil {
				return nil, err
			}
			if valid && done {
				result = append(result, operands)
			}
		default:
			return nil, errors.New("unsupported operation")
		}

		o, err = NextOperation(reader)
	}

	if !errors.Is(err, io.EOF) {
		return nil, err
	}

	return result, nil
}

func readOperands(r *bufio.Reader) (string /* valid */, bool /* done */, bool, error) {
	operands := ""

	c, _, err := r.ReadRune()
	for err == nil {
		if unicode.IsDigit(c) || c == ',' {
			operands += fmt.Sprintf("%c", c)
		} else if c == ')' {
			return operands, true, true, nil
		} else {
			return "", false, true, nil
		}
		c, _, err = r.ReadRune()
	}

	return "", false, false, err
}

func findIdx(a []string, f func(string) bool) int {
	for i, s := range a {
		if f(s) {
			return i
		}
	}
	return -1
}
