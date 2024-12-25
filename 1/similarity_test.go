package main_test

import (
	"fmt"
	"testing"

	aoc "github.com/antklim/aoc24/1"
)

func TestSimilarity(t *testing.T) {
	t.Run("returns similarity score", func(t *testing.T) {
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
				a:    []int{1, 2},
				b:    []int{},
				want: 0,
			},
			{
				a:    []int{},
				b:    []int{1, 2},
				want: 0,
			},
			{
				a:    []int{2},
				b:    []int{1, 2},
				want: 2,
			},
			{
				a:    []int{3, 4, 2, 1, 3, 3},
				b:    []int{4, 3, 5, 3, 9, 3},
				want: 31,
			},
		}

		for i, tC := range testCases {
			t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
				got := aoc.Similarity(tC.a, tC.b)

				if got != tC.want {
					t.Errorf("expected similarity score %d, got: %d", tC.want, got)
				}
			})
		}
	})
}
