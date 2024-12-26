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

func (r Report) IsSafeWithDampener() bool {
	for i := 0; i < len(r.levels); i++ {
		var a []int
		if i == 0 {
			a = r.levels[1:]
		} else if i == len(r.levels)-1 {
			a = r.levels[:len(r.levels)-1]
		} else {
			a = append(a, r.levels[:i]...)
			a = append(a, r.levels[i+1:]...)
		}

		if r := NewReport(a); r.IsSafe() {
			return true
		}
	}

	return false
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
