package quicksort

func quicksort(inputSlice []int, comparisons int) ([]int, int) {
	if len(inputSlice) < 2 {
		return inputSlice, 0
	}

	outputSlice := make([]int, len(inputSlice))
	copy(outputSlice, inputSlice)

	return outputSlice, len(inputSlice) - 1
}
