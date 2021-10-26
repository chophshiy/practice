package sets

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIsFlawless(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	cases := []struct {
		give   []int
		expect bool
	}{
		{[]int{1, 2, 3, 4}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 7, 9, 10}, false},
		{[]int{2, 2, 3, 4}, false},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case=%2d", i), func(t *testing.T) {
			rand.Shuffle(len(tt.give), func(i, j int) {
				tt.give[i], tt.give[j] = tt.give[j], tt.give[i]
			})
			got := IsFlawless(tt.give)
			if got != tt.expect {
				t.Errorf("wrong result: got %t, want %t", got, tt.expect)
			}
		})
	}
}
