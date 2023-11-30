package ex1

import "testing"

func TestSetIthBit(t *testing.T) {
	for _, c := range []struct {
		inNum, idx, bit, want int
	}{
		{2, 3, 1, 10},
		{3, 5, 0, 3},
	} {
		got := SetIthBit(c.inNum, c.idx, c.bit)
		if got != c.want {
			t.Errorf("SetIthBit(%d, %d, %d) == %d, want %d", c.inNum, c.idx, c.bit, got, c.want)
		}
	}
}
