package main_test

import (
	"fmt"
	"testing"

	aoc "github.com/antklim/aoc24/2"
)

func TestIsIncreasing(t *testing.T) {
	testCases := []struct {
		a    []int
		want bool
	}{
		{
			a:    nil,
			want: false,
		},
		{
			a:    []int{},
			want: false,
		},
		{
			a:    []int{1},
			want: false,
		},
		{
			a:    []int{1, 2},
			want: true,
		},
		{
			a:    []int{1, 1, 2, 4},
			want: true,
		},
		{
			a:    []int{1, 2, 2, 4},
			want: true,
		},
		{
			a:    []int{4, 2, 1},
			want: false,
		},
		{
			a:    []int{1, 2, 1, 4},
			want: false,
		},
		{
			a:    []int{1, 3, 5},
			want: true,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := aoc.IsIncreasing(tC.a)
			if got != tC.want {
				t.Errorf("expected is increasing to be %t, got: %t", tC.want, got)
			}
		})
	}
}

func TestIsDecreasing(t *testing.T) {
	testCases := []struct {
		a    []int
		want bool
	}{
		{
			a:    nil,
			want: false,
		},
		{
			a:    []int{},
			want: false,
		},
		{
			a:    []int{1},
			want: false,
		},
		{
			a:    []int{3, 2},
			want: true,
		},
		{
			a:    []int{3, 3, 2, 1},
			want: true,
		},
		{
			a:    []int{4, 2, 2, 1},
			want: true,
		},
		{
			a:    []int{1, 2, 4},
			want: false,
		},
		{
			a:    []int{4, 2, 1, 4},
			want: false,
		},
		{
			a:    []int{5, 3, 1},
			want: true,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := aoc.IsDecreasing(tC.a)
			if got != tC.want {
				t.Errorf("expected is decreasing to be %t, got: %t %v", tC.want, got, tC.a)
			}
		})
	}
}

func TestIsSafe(t *testing.T) {
	testCases := []struct {
		a    []int
		want bool
	}{
		{
			a:    nil,
			want: false,
		},
		{
			a:    []int{},
			want: false,
		},
		{
			a:    []int{1},
			want: false,
		},
		{
			a:    []int{3, 3, 2, 1},
			want: false,
		},
		{
			a:    []int{4, 2, 2, 1},
			want: false,
		},
		{
			a:    []int{1, 2, 2, 4},
			want: false,
		},
		{
			a:    []int{4, 2, 1, 4},
			want: true,
		},
		{
			a:    []int{7, 6, 4, 2, 1},
			want: true,
		},
		{
			a:    []int{1, 2, 7, 8, 9},
			want: false,
		},
		{
			a:    []int{9, 7, 6, 2, 1},
			want: false,
		},
		{
			a:    []int{1, 3, 2, 4, 5},
			want: true,
		},
		{
			a:    []int{8, 6, 4, 4, 1},
			want: false,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := aoc.IsSafe(tC.a)
			if got != tC.want {
				t.Errorf("expected is safe to be %t, got: %t %v", tC.want, got, tC.a)
			}
		})
	}
}

func TestReportIsSafe(t *testing.T) {
	testCases := []struct {
		a    []int
		want bool
	}{
		{
			a:    []int{7, 6, 4, 2, 1},
			want: true,
		},
		{
			a:    []int{1, 2, 7, 8, 9},
			want: false,
		},
		{
			a:    []int{9, 7, 6, 2, 1},
			want: false,
		},
		{
			a:    []int{1, 3, 2, 4, 5},
			want: false,
		},
		{
			a:    []int{8, 6, 4, 4, 1},
			want: false,
		},
		{
			a:    []int{1, 3, 6, 7, 9},
			want: true,
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			report := aoc.NewReport(tC.a)
			got := report.IsSafe()
			if got != tC.want {
				t.Errorf("expected report is safe to be %t, got: %t %v", tC.want, got, tC.a)
			}
		})
	}
}
