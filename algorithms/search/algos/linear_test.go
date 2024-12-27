package algos

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	type TestCase struct {
		Array         []int
		Value         int
		ExpectedFound bool
		ExpectedIndex int
	}

	testData := []TestCase{
		{
			Array:         []int{2, 4, 6, 8, 10},
			Value:         8,
			ExpectedFound: true,
			ExpectedIndex: 3,
		},
		{
			Array:         []int{5, 10, 15, 21, 37, 100},
			Value:         2,
			ExpectedFound: false,
			ExpectedIndex: -1,
		},
	}

	for _, testCase := range testData {
		testName := fmt.Sprintf("Search for %d in %v", testCase.Value, testCase.Array)
		t.Run(testName, func(t *testing.T) {
			found, index := LinearSearch(testCase.Array, testCase.Value)
			if !(found == testCase.ExpectedFound && index == testCase.ExpectedIndex) {
				t.Errorf("Expected found=%t, index=%d. Got found=%t, index=%d.",
					testCase.ExpectedFound, testCase.ExpectedIndex, found, index)
			}
		})
	}
}
