package main

func Similarity(a, b []int) int {
	bScore := score(b)

	result := 0
	for _, v := range a {
		result += v * bScore[v]
	}
	return result
}

func score(a []int) map[int]int {
	result := make(map[int]int)

	for _, v := range a {
		score := result[v]
		result[v] = score + 1
	}

	return result
}
