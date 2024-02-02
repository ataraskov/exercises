package armstrong

import (
	"math"
)

func IsNumber(n int) bool {
	if n < 0 {
		return false
	}

	res := n
	nums := []int{}
	for i := 1; res > 0; i++ {
		nums = append(nums, res%10)
		res = res / 10
	}

	res = 0
	for _, v := range nums {
		res += (int)(math.Pow((float64)(v), (float64)(len(nums))))
	}
	return res == n
}
