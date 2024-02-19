package main

import (
	"fmt"

	// Import the algos package
	"github.com/donovan-said/go-examples/internal/algos"
)

func main() {
	fmt.Println(">> SORTING ALGORITHMS")

	mySliceInts := []int{5, 2, 6, 3, 1, 4}
	mySliceWords := []string{"banana", "apple", "cucumber"}
	n, w := algos.Sorting(mySliceInts, mySliceWords)
	fmt.Printf("Sorted numbers are: %v\n", n)
	fmt.Printf("Sorted works are: %v", w)
}
