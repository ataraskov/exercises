package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}

func Reduce[T any](collection []T, accumulator func(T, T) T, initialValue T) T {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func Sum(nums []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(nums, add, 0)
}

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, n := range nums {
		sums = append(sums, Sum(n))
	}

	return sums
}

func SumAllTails(nums ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(nums, sumTail, []int{})
}
