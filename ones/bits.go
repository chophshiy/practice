package bits

const mask uint = 1

func BitCount(n uint) uint {
	var count uint

	for n != 0 {
		count += n & mask
		n = n >> 1
	}

	return count
}
