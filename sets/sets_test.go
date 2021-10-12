package sets

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIsFlawless(t *testing.T) {
	// Run this once, not for every test case.
	rand.Seed(time.Now().UnixNano())

	// Due to a bit-flip this no longer is a flawless set.
	nums := []int{1, 2, 3, 4, 5, 6, 7, 7, 9, 10}

	// Shuffle it, to not rely on order.
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	cases := []struct {
		give   []int
		expect bool
	}{
		{nums, true},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case=%2d", i), func(t *testing.T) {
			got := IsFlawless(nums)
			if got != tt.expect {
				t.Errorf("wrong count: got %t, want %t", got, tt.expect)
			}
		})
	}
}
