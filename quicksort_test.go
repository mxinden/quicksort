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
	{filePath: "./examples/1", expectedComparisons: 0},
	{filePath: "./examples/empty", expectedComparisons: 0},
}

func TestQuickSort(t *testing.T) {
	for _, test := range quickSortTests {
		inputSlice, err := fileToIntSlice(test.filePath)
		if err != nil {
			t.Fatal(err)
		}
		outputSlice, amountComparisons := quicksort(inputSlice, 0)

		sortedSlice := make([]int, len(inputSlice))
		copy(sortedSlice, inputSlice)
		sort.Ints(sortedSlice)

		if !reflect.DeepEqual(outputSlice, sortedSlice) {
			t.Fatal(fmt.Sprintf("output slice not sorted, expected %v, got %v", sortedSlice, outputSlice))
		}

		if test.expectedComparisons != amountComparisons {
			t.Fatal(fmt.Sprintf("amountComparisons does not equal expected comparisons, expected %v, got %v", test.expectedComparisons, amountComparisons))
		}

	}
}
