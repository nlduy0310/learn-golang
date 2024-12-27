package algos

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type TestCase struct {
		Array         []int
		Value         int
		ExpectedFound bool
		ExpectedIndex int
	}

	testData := []TestCase{
		{
			Array:         []int{2, 4, 6, 8, 10},
			Value:         7,
			ExpectedFound: false,
			ExpectedIndex: -1,
		},
		{
			Array:         []int{5, 10, 15, 20, 25},
			Value:         10,
			ExpectedFound: true,
			ExpectedIndex: 1,
		},
		{
			Array:         []int{3, 6, 9, 12},
			Value:         5,
			ExpectedFound: false,
			ExpectedIndex: -1,
		},
	}

	for _, testCase := range testData {
		testName := fmt.Sprintf("Search for %d in %v", testCase.Value, testCase.Array)
		t.Run(testName, func(ct *testing.T) {
			found, index := BinarySearch(testCase.Array, testCase.Value)
			if !(found == testCase.ExpectedFound && index == testCase.ExpectedIndex) {
				ct.Errorf("Expected found=%t, index=%d. Got found=%t, index=%d",
					testCase.ExpectedFound, testCase.ExpectedIndex, found, index)
			}
		})
	}
}
