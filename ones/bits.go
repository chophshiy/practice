package bits

const mask uint = 1 << 63

func BitCount(n uint) int {
	var count int
	b := n

	for i := 0; i < 64; i++ {
		if b&mask == mask {
			count += 1
		}
		b = b << 1
	}

	return count
}
