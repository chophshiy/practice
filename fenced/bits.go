package bits

const mask uint64 = 1

func MaxFence(n uint64) int {
	var maxCount, cur int
	var leadingOne, trailingOne bool

	for n != 0 {
		if n&mask == mask {
			if !leadingOne {
				leadingOne = true
			} else {
				trailingOne = true
			}
			if leadingOne && trailingOne {
				if cur > maxCount {
					maxCount = cur
					cur = 0
					trailingOne = false
				}
			}
		} else {
			if leadingOne {
				cur += 1
			}
		}
		n = n >> 1
	}

	return maxCount
}
