package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

func ReadMul(r io.Reader) ([]string, error) {
	reader := bufio.NewReader(r)

	operation := ""
	operands := ""

	var result []string

	c, _, err := reader.ReadRune()
	for err == nil {
		if operation == "" && c == 'm' {
			operation = "m"
		} else if operation == "m" {
			if c == 'u' {
				operation = "mu"
			} else {
				operation = ""
			}
		} else if operation == "mu" {
			if c == 'l' {
				operation = "mul"
			} else {
				operation = ""
			}
		} else if operation == "mul" {
			if c == '(' {
				operation = "mul("
			} else {
				operation = ""
			}
		} else if operation == "mul(" {
			if unicode.IsDigit(c) || c == ',' {
				operands += fmt.Sprintf("%c", c)
			} else if c == ')' {
				result = append(result, operands)
				operands = ""
				operation = ""
			} else {
				operands = ""
				operation = ""
			}
		}

		c, _, err = reader.ReadRune()
	}

	if err == io.EOF {
		err = nil
	}

	return result, err
}
