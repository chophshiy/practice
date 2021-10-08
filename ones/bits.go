package bits

const mask uint = 1

func BitCount(n uint) uint {
	var count uint
	b := n

	for i := 0; i < 64; i++ {
		count += b & mask
		b = b >> 1
	}

	return count
}
