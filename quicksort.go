package quicksort

func quicksort(inputSlice []int, f preparePivot) int {
	if len(inputSlice) < 2 {
		return 0
	}

	return s(inputSlice, 0, len(inputSlice)-1, f)
}

func s(slice []int, start, end int, f preparePivot) int {
	if end-start < 1 {
		return 0
	}
	f(slice, start, end)
	pivot := slice[start]
	small := start

	for i := start + 1; i <= end; i = i + 1 {
		if slice[i] < pivot {
			swap(slice, i, small+1)
			small = small + 1
		}
	}

	swap(slice, start, small)

	a, b := 0, 0
	a = s(slice, start, small-1, f)
	b = s(slice, small+1, end, f)

	return a + b + (end - start)
}

func swap(slice []int, a, b int) {
	tmp := slice[a]
	slice[a] = slice[b]
	slice[b] = tmp
}

type preparePivot func(slice []int, start, end int)

func firstElementPivot(slice []int, start, end int) {}

func finalElementPivot(slice []int, start, end int) {
	swap(slice, start, end)
}

func medianPivot(slice []int, start, end int) {
	if end-start < 2 {
		return
	}

	middleIndex := start + (end-start)/2

	first := slice[start]
	middle := slice[middleIndex]
	last := slice[end]

	if (first < middle && middle < last) || (last < middle && middle < first) {
		swap(slice, start, middleIndex)
	} else if (first < last && last < middle) || (middle < last && last < first) {
		swap(slice, start, end)
	}
}
