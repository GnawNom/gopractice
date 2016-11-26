package sort

func BubbleSort(data Interface) {
	n := data.Len()
	for {
		swapped := false
		for i := 0; i < n-1; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
