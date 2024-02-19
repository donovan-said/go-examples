package main

import (
	"fmt"

	// Import the algos package
	"github.com/donovan-said/go-examples/internal/algos"
)

func main() {
	fmt.Println(">> SORTING ALGORITHMS")

	mySliceInts := []int{5, 2, 6, 3, 1, 4}
	algos.SortNums(mySliceInts)

	mySliceWords := []string{"banana", "apple", "cucumber"}
	algos.SortWords(mySliceWords)

}
