package strings

import (
	"fmt"
	"testing"
)

func TestRuneLength(t *testing.T) {
	cases := []struct {
		give   string
		expect uint
	}{
		{"abc", 3},
		{" def", 4},
		{"牛奶", 2},
		{"牛奶.", 3},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case=%2d", i), func(t *testing.T) {
			got := RuneLength(tt.give)
			if got != tt.expect {
				t.Errorf("wrong length: got %d, want %d", got, tt.expect)
			}
		})
	}
}
