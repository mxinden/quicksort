package quicksort

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func fileToIntSlice(filePath string) ([]int, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []int{}, err
	}

	stringLines := strings.Split(string(content), "\n")

	intLines, err := StringSliceToIntSlice(stringLines)
	if err != nil {
		return []int{}, err
	}

	return intLines, nil
}

func StringSliceToIntSlice(stringLines []string) ([]int, error) {
	intSlice := []int{}
	for _, s := range stringLines {
		// strings.Split adds an empty string at the end of the slice.
		if s == "" {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return intSlice, err
		}
		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}
