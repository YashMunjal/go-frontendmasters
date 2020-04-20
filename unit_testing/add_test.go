package unit_testing_test

import "testing"

func testAdd(t *testing.T) {
	expected := 4
	actual := addNums(2, 2)

	if actual != expected {
		t.Errorf("Doesn't match up")
	}
}
