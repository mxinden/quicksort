package quicksort

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var quickSortTests = []struct {
	filePath    string
	comparisons []int
}{
	{filePath: "./examples/empty", comparisons: []int{0, 0, 0}},
	{filePath: "./examples/1", comparisons: []int{0, 0, 0}},
	{filePath: "./examples/12", comparisons: []int{1, 1, 1}},
	{filePath: "./examples/21", comparisons: []int{1, 1, 1}},
	{filePath: "./examples/123", comparisons: []int{3, 3, 2}},
	{filePath: "./examples/132", comparisons: []int{3, 2, 2}},
	{filePath: "./examples/213", comparisons: []int{2, 3, 2}},
	{filePath: "./examples/231", comparisons: []int{2, 3, 2}},
	{filePath: "./examples/312", comparisons: []int{3, 2, 2}},
	{filePath: "./examples/321", comparisons: []int{3, 3, 2}},
	{filePath: "./examples/length-5", comparisons: []int{6, 10, 6}},
	{filePath: "./examples/length-10", comparisons: []int{26, 21, 21}},
	{filePath: "./examples/length-20", comparisons: []int{69, 65, 56}},
	{filePath: "./examples/length-1000000", comparisons: []int{24308571, 25371973, 21214423}},
}

var preparePivotFunctions = []preparePivot{firstElementPivot, finalElementPivot, medianPivot}

func TestQuickSort(t *testing.T) {
	for _, test := range quickSortTests {
		for i, f := range preparePivotFunctions {
			inputSlice, err := fileToIntSlice(test.filePath)
			if err != nil {
				t.Fatal(err)
			}

			outputSlice := make([]int, len(inputSlice))
			copy(outputSlice, inputSlice)

			amountComparisons := quicksort(outputSlice, f)

			sortedSlice := make([]int, len(inputSlice))
			copy(sortedSlice, inputSlice)
			sort.Ints(sortedSlice)

			if !reflect.DeepEqual(outputSlice, sortedSlice) {
				t.Fatal(fmt.Sprintf("%v: output slice not sorted, expected %v, got %v", test.filePath, sortedSlice, outputSlice))
			}

			if test.comparisons[i] != amountComparisons {
				t.Fatal(fmt.Sprintf("%v: amount comparisons does not equal expected comparisons, expected %v, got %v", test.filePath, test.comparisons[i], amountComparisons))
			}

		}

	}
}

var medianPivotTests = []struct {
	input  []int
	output []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{1, 2}},
	{[]int{2, 1}, []int{2, 1}},
	{[]int{1, 2, 3}, []int{2, 1, 3}},
	{[]int{1, 2, 3, 4}, []int{2, 1, 3, 4}},
	{[]int{8, 2, 4, 5, 7, 1}, []int{4, 2, 8, 5, 7, 1}},
}

func TestMedianPivot(t *testing.T) {
	for _, test := range medianPivotTests {
		outputSlice := make([]int, len(test.input))
		copy(outputSlice, test.input)
		medianPivot(outputSlice, 0, len(outputSlice)-1)

		if !reflect.DeepEqual(outputSlice, test.output) {
			t.Fatal(fmt.Sprintf("expected %v, got %v", test.output, outputSlice))
		}
	}
}
