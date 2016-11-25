package sort

import "testing"

func TestIsSorted(t *testing.T) {
	var tests = []struct {
		numbers IntSlice
		sorted  bool
	}{
		{[]int{1, 2, 3}, true},
		{[]int{3, 2, 1}, false},
		{[]int{0, 0, 0}, true},
	}
	for _, c := range tests {
		if IsSorted(c.numbers) != c.sorted {
			t.Errorf("IsSorted failed on %q", c.numbers)
		}
	}
}
