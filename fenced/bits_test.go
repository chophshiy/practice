package bits

import (
	"fmt"
	"testing"
)

func TestMaxFence(t *testing.T) {
	cases := []struct {
		give   uint
		expect int
	}{
		{0b0001000, 0},
		{0b11, 0},
		{0b1001, 2},
		{0b110001001, 3},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case=%2d", i), func(t *testing.T) {
			got := MaxFence(tt.give)
			if got != tt.expect {
				t.Errorf("wrong count: got %d, want %d", got, tt.expect)
			}
		})
	}
}
