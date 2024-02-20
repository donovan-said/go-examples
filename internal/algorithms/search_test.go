package algorithms

import (
	"testing"
)

func TestSearch(t *testing.T) {

	test_input := []int{1, 2, 3, 4, 5, 6}
	expected_output := 3
	actual_output := Search(test_input)

	if actual_output != expected_output {
		t.Errorf("Expected %v but got %v", expected_output, actual_output)
	}

}
