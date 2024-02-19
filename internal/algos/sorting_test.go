package algos

import (
	"testing"
)

func TestSortNums(t *testing.T) {

	mySliceInts := []int{5, 2, 6, 3, 1, 4}
	ordered_num := SortNums(mySliceInts)
	expected_slice := []int{1, 2, 3, 4, 5, 6}

	for k, v := range ordered_num {
		if v != expected_slice[k] {
			t.Errorf("Mismatch in slice!")
		}
	}
}
