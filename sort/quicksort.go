package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func randInt(i int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(i)
}

func QuickSort(data Interface) {
	quickSort(data, 0, data.Len()-1)
}

func partition(data Interface, start, end int) int {
	// One element arrays are already sorted
	fmt.Printf("pivot called with %v %v\n", start, end)
	length := end - start + 1
	if length == 1 {
		return start
	}
	pivot := start + randInt(length)
	data.Swap(pivot, end)
	for i := start; i < end; i++ {
		// data[i] <= data[end]
		if data.Less(i, end) {
			data.Swap(i, start)
			start++
		}
	}
	data.Swap(start, end)
	return start
}

func quickSort(data Interface, start, end int) {
	if start < end {
		pivot := partition(data, start, end)
		quickSort(data, start, pivot-1)
		quickSort(data, pivot+1, end)
	}
}
