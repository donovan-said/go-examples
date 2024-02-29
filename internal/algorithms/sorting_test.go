package algorithms

import (
	"testing"
)

func TestSortNums(t *testing.T) {

	test_input := []int{5, 2, 6, 3, 1, 4}
	expected_output := []int{1, 2, 3, 4, 5, 6}
	actual_output := SortNums(test_input)

	for k, v := range actual_output {
		if v != expected_output[k] {
			t.Errorf("Mismatch in slice!")
		}
	}
}

func TestSortWords(t *testing.T) {

	test_input := []string{"banana", "apple", "cucumber"}
	expected_output := []string{"apple", "banana", "cucumber"}
	actual_output := SortWords(test_input)

	for k, v := range actual_output {
		if v != expected_output[k] {
			t.Errorf("Mismatch in slice!")
		}
	}
}
