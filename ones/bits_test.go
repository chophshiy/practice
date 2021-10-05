package bits

import (
	"fmt"
	"testing"
)

func TestBitCount(t *testing.T) {
	cases := []struct {
		give   uint
		expect int
	}{
		{4294967296, 1}, // These numbers should be just outside of the uint32 range.
		{4294967297, 2}, // Conversion of case two should be: 1 00000000 00000000 00000000 00000001
		{4294967298, 2},
		{4294967299, 3},
		{4294967300, 2},
		{4294967301, 3},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case=%2d", i), func(t *testing.T) {
			got := BitCount(tt.give)
			if got != tt.expect {
				t.Errorf("wrong count: got %d, want %d", got, tt.expect)
			}
		})
	}
}
