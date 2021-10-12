package sets

func IsFlawless(nums []int) bool {
	var compare []int
	var total int
	var result bool = true

	for place, _ := range compare {
		compare[place] = place + 1
	}

	for outer, _ := range nums {
		for inner, _ := range compare {
			if compare[inner]^nums[outer] == 0 {
				compare[inner] = 0
			}
		}
	}

	for place, _ := range compare {
		total += compare[place]
	}

	if total != 0 {
		result = false
	}

	return result
}
