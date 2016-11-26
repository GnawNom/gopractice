package sort

import "testing"

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

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

func TestInsertionSort(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	InsertionSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("got    %v", data)
	}
}

func TestBubbleSort(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	BubbleSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("got    %v", data)
	}
}

func TestQuickSort(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	QuickSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("got    %v", data)
	}
}
