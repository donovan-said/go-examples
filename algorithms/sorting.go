package main

import (
	"fmt"
	"sort"
)

func sorting_algo() {
	/*
		Sorting integers:

		The following sorts a slice of integers using quicksort
	*/
	unsorted_nums := []int{5, 2, 6, 3, 1, 4}
	fmt.Printf("Unsorted numbers: %v\n", unsorted_nums)
	sort.Ints(unsorted_nums)
	fmt.Printf("Sorted numbers: %v\n", unsorted_nums)

	/*
		Sorting strings:

		The following sorts a slice of strings using mergesort
	*/
	unsorted_words := []string{"banana", "apple", "cucumber"}
	fmt.Printf("Unsorted words: %v\n", unsorted_words)
	sort.Strings(unsorted_words)
	fmt.Printf("Sorted works: %v\n", unsorted_words)
}
