package main

import (
	"fmt"

	// Import the algos package
	"github.com/donovan-said/go-examples/internal/algos"
)

func algorithms() {
	// Sorting Algorithms
	fmt.Println(">> SORTING ALGORITHMS")

	SortSliceInts := []int{5, 2, 6, 3, 1, 4}
	algos.SortNums(SortSliceInts)

	SortSliceWords := []string{"banana", "apple", "cucumber"}
	algos.SortWords(SortSliceWords)

	// Sorting Algorithms
	fmt.Println(">> SEARCH ALGORITHMS")

	SearchSliceInts := []int{1, 2, 3, 4, 5, 6}
	algos.Search(SearchSliceInts)
}

func main() {

	algorithms()
}
