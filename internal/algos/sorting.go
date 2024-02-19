package algos

import (
	"fmt"
	"sort"
)

func Sorting(nums []int, words []string) (ordered_num []int, ordered_words []string) {
	/*
		Sorting integers:

		The following sorts a slice of integers using quicksort
	*/

	fmt.Printf("Received unsorted numbers: %v\n", nums)
	sort.Ints(nums)

	/*
		Sorting strings:

		The following sorts a slice of strings using mergesort
	*/

	fmt.Printf("Received unsorted words: %v\n", words)
	sort.Strings(words)

	return nums, words
}
