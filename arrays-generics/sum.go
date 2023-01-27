package main

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func Find[T any](items []T, predicate func(T) bool) (value T, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
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
