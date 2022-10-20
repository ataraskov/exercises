package main

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, n := range nums {
		sums = append(sums, Sum(n))
	}

	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int

	for _, n := range nums {
		if len(n) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(n[1:]))
		}
	}

	return sums
}
