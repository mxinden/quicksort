package quicksort

func quicksort(inputSlice []int) int {
	if len(inputSlice) < 2 {
		return 0
	}

	return s(inputSlice, 0, len(inputSlice)-1)
}

func s(slice []int, start, end int) int {
	if end-start < 1 {
		return 0
	}
	pivot := slice[start]
	small := start

	for i := start + 1; i <= end; i = i + 1 {
		if slice[i] < pivot {
			swap(slice, i, small+1)
			small = i
		}
	}

	a, b := 0, 0
	a = s(slice, start+1, small)
	b = s(slice, small+1, end)

	swap(slice, start, small)

	return a + b + (end - start)
}

func swap(slice []int, a, b int) {
	tmp := slice[a]
	slice[a] = slice[b]
	slice[b] = tmp
}
