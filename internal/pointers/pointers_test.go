package pointers

import "testing"

func TestModifyingPointers(t *testing.T) {
	test_input_1 := "shark"
	test_input_2 := "jellyfish"
	expected_output := "jellyfish"
	actual_output := ModifyingPointers(test_input_1, test_input_2)

	if expected_output != actual_output {
		t.Errorf("Expected output %v, but got %v", expected_output, actual_output)
	}
}
