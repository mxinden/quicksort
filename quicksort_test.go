package quicksort

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var quickSortTests = []struct {
	filePath            string
	expectedComparisons int
}{
	{filePath: "./examples/empty", expectedComparisons: 0},
	{filePath: "./examples/1", expectedComparisons: 0},
	{filePath: "./examples/12", expectedComparisons: 1},
	{filePath: "./examples/21", expectedComparisons: 1},
	{filePath: "./examples/123", expectedComparisons: 3},
	{filePath: "./examples/132", expectedComparisons: 3},
	{filePath: "./examples/213", expectedComparisons: 2},
	{filePath: "./examples/231", expectedComparisons: 2},
	{filePath: "./examples/312", expectedComparisons: 3},
	{filePath: "./examples/321", expectedComparisons: 3},
	{filePath: "./examples/length-1000000", expectedComparisons: 24308571},
}

func TestQuickSort(t *testing.T) {
	for _, test := range quickSortTests {
		inputSlice, err := fileToIntSlice(test.filePath)
		if err != nil {
			t.Fatal(err)
		}

		outputSlice := make([]int, len(inputSlice))
		copy(outputSlice, inputSlice)

		amountComparisons := quicksort(outputSlice, firstElementPivot)

		sortedSlice := make([]int, len(inputSlice))
		copy(sortedSlice, inputSlice)
		sort.Ints(sortedSlice)

		if !reflect.DeepEqual(outputSlice, sortedSlice) {
			t.Fatal(fmt.Sprintf("%v: output slice not sorted, expected %v, got %v", test.filePath, sortedSlice, outputSlice))
		}

		if test.expectedComparisons != amountComparisons {
			t.Fatal(fmt.Sprintf("%v: amount comparisons does not equal expected comparisons, expected %v, got %v", test.filePath, test.expectedComparisons, amountComparisons))
		}
	}
}

var medianPivotTests = []struct {
	input  []int
	output []int
}{
	{[]int{}, []int{}},
}

func TestMedianPivot(t *testing.T) {
	for _, test := range medianPivotTests {
		outputSlice := make([]int, len(test.input))
		copy(outputSlice, test.input)
		medianPivot(outputSlice, 0, len(outputSlice))

		if !reflect.DeepEqual(outputSlice, test.output) {
			t.Fatal(fmt.Sprintf("expected %v, got %v", test.output, outputSlice))
		}
	}
}
