package main

import (
	"fmt"

	// Import the local package
	"github.com/donovan-said/go-examples/internal/algorithms"
	"github.com/donovan-said/go-examples/internal/pointers"
)

func algos() {
	// Sorting Algorithms
	fmt.Println(">> SORTING ALGORITHMS")

	SortSliceInts := []int{5, 2, 6, 3, 1, 4}
	algorithms.SortNums(SortSliceInts)

	SortSliceWords := []string{"banana", "apple", "cucumber"}
	algorithms.SortWords(SortSliceWords)

	// Sorting Algorithms
	fmt.Println(">> SEARCH ALGORITHMS")

	SearchSliceInts := []int{1, 2, 3, 4, 5, 6}
	algorithms.Search(SearchSliceInts)
}

func points() {
	fmt.Println(">> POINTER EXAMPLES")

	pointers.UsingPointers()
}

func main() {

	algos()
	points()
}
