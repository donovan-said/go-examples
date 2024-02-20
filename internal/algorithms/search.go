package algorithms

import (
	"fmt"
	"sort"
)

func Search(search []int) (result int) {
	/* Search for the index of an element in a sorted slice of integers using
	binary search */

	index := sort.SearchInts(search, 4)
	fmt.Printf("Index placement is: %v\n", index)

	return index
}
