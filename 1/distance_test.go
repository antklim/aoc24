package main_test

import (
	"errors"
	"fmt"
	"testing"

	aoc "github.com/antklim/aoc24/1"
)

func TestDistance(t *testing.T) {
	t.Run("returns error when lists have different lengths", func(t *testing.T) {
		want := aoc.ErrDifferentLength
		_, err := aoc.Distance([]int{}, []int{1})
		if !errors.Is(err, want) {
			t.Errorf("expected error: %s, but got: %s", want, err)
		}
	})

	t.Run("returns distance between lists", func(t *testing.T) {
		testCases := []struct {
			a    []int
			b    []int
			want int
		}{
			{
				a:    nil,
				b:    nil,
				want: 0,
			},
			{
				a:    []int{},
				b:    []int{},
				want: 0,
			},
			{
				a:    []int{3, 4, 2, 1, 3, 3},
				b:    []int{4, 3, 5, 3, 9, 3},
				want: 11,
			},
		}

		for i, tC := range testCases {
			t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
				got, err := aoc.Distance(tC.a, tC.b)
				if err != nil {
					t.Fatalf("got unexpected error: %s", err)
				}

				if got != tC.want {
					t.Errorf("expected distance %d, got: %d", tC.want, got)
				}
			})
		}
	})
}
