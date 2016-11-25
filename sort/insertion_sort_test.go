package sort

import "testing"

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestInsertionSort(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	InsertionSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("got    %v", data)
	}
}
