package main

type Direction int

const (
	Unknown    Direction = 0
	Increasing Direction = 1
	Decreasing Direction = -1
)

type Report struct {
	levels    []int
	direction Direction
}

func NewReport(a []int) Report {
	r := Report{
		levels:    a,
		direction: Unknown,
	}

	if IsIncreasing(a) {
		r.direction = Increasing
	}
	if IsDecreasing(a) {
		r.direction = Decreasing
	}

	return r
}

func (r Report) IsSafe() bool {
	if r.direction == Unknown {
		return false
	}

	return IsSafe(r.levels)
}

func IsSafe(a []int) bool {
	if len(a) < 2 {
		return false
	}

	for i, v := range a[1:] {
		diff := abs(v - a[i])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func IsIncreasing(a []int) bool {
	if len(a) < 2 {
		return false
	}

	for i, v := range a[1:] {
		if v < a[i] {
			return false
		}
	}

	return true
}

func IsDecreasing(a []int) bool {
	if len(a) < 2 {
		return false
	}

	for i, v := range a[1:] {
		if v > a[i] {
			return false
		}
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
