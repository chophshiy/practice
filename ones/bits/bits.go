package bits

const mask = 1 << 63

func BitCount(n uint) int {
	var count int
	b := n

	for i := 0; i < 64; i++ {
		if b<<1&mask == mask {
			count += 1
		}
	}

	return count
}
