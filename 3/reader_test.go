package main_test

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	aoc "github.com/antklim/aoc24/3"
)

func TestReadMul(t *testing.T) {
	testCases := []struct {
		s    string
		want []string
	}{
		{
			s:    "",
			want: nil,
		},
		{
			s:    "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			want: []string{"2,4", "5,5", "11,8", "8,5"},
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			buf := strings.NewReader(tC.s)

			got, err := aoc.ReadMul(buf)

			if err != nil {
				t.Fatalf("expected no read errors, but got: %s", err)
			}

			if w, g := len(tC.want), len(got); w != g {
				t.Logf("want: %v\n got: %v", tC.want, got)
				t.Fatalf("invalid amount of operations, want: %d, got: %d", w, g)
			}

			var errs error
			for i, o := range got {
				if tC.want[i] != o {
					errs = errors.Join(errs, fmt.Errorf("operation #%d is not equal, want: %s, got: %s", i, tC.want[i], o))
				}
			}

			if errs != nil {
				t.Fatalf("operations are not equal: %s", errs)
			}
		})
	}
}

func TestNextOperation(t *testing.T) {
	s := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	buf := strings.NewReader(s)
	reader := bufio.NewReader(buf)

	want := []aoc.Operation{aoc.OperationMul, aoc.OperationDont, aoc.OperationMul, aoc.OperationMul, aoc.OperationMul, aoc.OperationDo, aoc.OperationMul}
	var got []aoc.Operation

	o, err := aoc.NextOperation(reader)
	for err == nil {
		got = append(got, o)
		o, err = aoc.NextOperation(reader)
	}
	if !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}

	if lw, lg := len(want), len(got); lw != lg {
		t.Logf("want operations: %v", want)
		t.Logf("got operations: %v", got)
		t.Fatalf("number of operations is not valid, want: %d, got %d", lw, lg)
	}

	var errs error
	for i, w := range want {
		if w != got[i] {
			errs = errors.Join(errs, fmt.Errorf("operation #%d is not valid, want: %d, got: %d", i, w, got[i]))
		}
	}

	if errs != nil {
		t.Fatalf("operations are not equal: %s", errs)
	}
}
