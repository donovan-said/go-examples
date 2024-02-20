package algorithms

import (
	"fmt"
	"sort"
)

func SortNums(nums []int) (ordered_num []int) {
	/*
		Sorting integers:

		The following sorts a slice of integers using quicksort
	*/

	fmt.Printf("Received unsorted numbers: %v\n", nums)
	sort.Ints(nums)
	fmt.Printf("Sorted numbers are: %v\n", nums)

	return nums
}

func SortWords(words []string) (ordered_words []string) {
	/*
		Sorting strings:

		The following sorts a slice of strings using mergesort
	*/

	fmt.Printf("Received unsorted words: %v\n", words)
	sort.Strings(words)
	fmt.Printf("Sorted words are: %v\n", words)

	return words
}
