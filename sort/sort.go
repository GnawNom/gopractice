package sort

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type SortFunction func(Interface)

type IntSlice []int

func (p IntSlice) Len() int { return len(p) }

func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(data Interface, sortFunction SortFunction) {
	sortFunction(data)
}

func IsSorted(data Interface) bool {
	for i := data.Len() - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
