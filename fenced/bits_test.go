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
		{0001000, 0},
		{11, 0},
		{1001, 2},
		{110001001, 3},
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
