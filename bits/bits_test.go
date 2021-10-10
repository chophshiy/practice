package bits

import (
	"fmt"
	"testing"
)

func TestBitCount(t *testing.T) {
	cases := []struct {
		give   uint
		expect uint
	}{
		{1, 1},
		{3, 2},
		{10, 2},
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

func TestMaxFence(t *testing.T) {
	cases := []struct {
		give   uint64
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
