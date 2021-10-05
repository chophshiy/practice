package bits

func BitCount(n uint) int {
	var count int
	b := n

	for i := 0; i < 64; i++ {
		// This fails all test cases by -1 if either the mask is 8 bits, or the mask below is expressed in hex.  I'm not understanding why.
		if b<<i&0b1000000000000000000000000000000000000000000000000000000000000000 == 0b1000000000000000000000000000000000000000000000000000000000000000 {
			count += 1
		}
	}

	return count
}
