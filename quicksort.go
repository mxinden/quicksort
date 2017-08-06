package quicksort

func quicksort(inputSlice []int, comparisons int) ([]int, int) {
	outputSlice := make([]int, len(inputSlice))
	copy(outputSlice, inputSlice)

	return outputSlice, comparisons
}
