package bits

const mask uint = 1

func MaxFence(n uint) int {
	var maxCount, cur int
	var leadingOne, trailingOne bool
	b := n

	for i := 0; i < 64; i++ {
		if b&mask == mask {
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
		b = b >> 1
	}

	return maxCount
}
