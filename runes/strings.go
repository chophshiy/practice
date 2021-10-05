package strings

func RuneLength(s string) uint {
	var count uint
	var i, stride int

	if len(s) == 0 {
		return 0
	}

	e := []byte(s) // To not rely on the string-to-rune conversion.

	for i < len(e) {
		b := e[i]
		switch {
		case b&0b11111100 == 0b11111100:
			stride = 6
		case b&0b11111000 == 0b11111000:
			stride = 5
		case b&0b11110000 == 0b11110000:
			stride = 4
		case b&0b11100000 == 0b11100000:
			stride = 3
		case b&0b11000000 == 0b11000000:
			stride = 2
		default:
			stride = 1
		}

		i += stride
		count += 1
	}

	return count
}
