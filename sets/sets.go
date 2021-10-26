package sets

func IsFlawless(nums []int) bool {
	var result bool = true
	var x int

	for i, e := range nums {
		x ^= e ^ (i + 1)
	}
	if x != 0 {
		result = false
	}

	return result
}
