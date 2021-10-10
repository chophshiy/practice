package bits

func BitCount(n uint) uint {
	var count uint

	for n > 0 {
		count += n & 0x01
		n >>= 1
	}

	return count
}

func MaxFence(n uint64) int {
	var maxCount, cur int
	var inFence bool

	for n > 0 {
		cur = 0
		for n&0x01 != 0x01 {
			cur += 1
			if n > 0 {
				n >>= 1
			}
		}
		if inFence {
			if cur > maxCount {
				maxCount = cur
			}
		} else {
			inFence = true
		}
		n >>= 1
	}

	return maxCount
}
