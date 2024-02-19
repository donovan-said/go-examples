package main

import (
	"fmt"

	// Import the algos package
	"github.com/donovan-said/go-examples/internal/algos"
)

func main() {
	fmt.Println(">> SORTING ALGORITHMS")

	mySliceInts := []int{5, 2, 6, 3, 1, 4}
	ordered_num := algos.SortNums(mySliceInts)
	fmt.Printf("Sorted numbers are: %v\n", ordered_num)

	mySliceWords := []string{"banana", "apple", "cucumber"}
	ordered_words := algos.SortWords(mySliceWords)
	fmt.Printf("Sorted works are: %v", ordered_words)
}
