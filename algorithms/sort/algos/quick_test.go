package algos

import (
	"fmt"
	"testing"
)

func TestQuickSortNaive(t *testing.T) {
	testData := []testCase{
		{
			Input:    []int{9, 7, 4, 2, 0, -3},
			Expected: []int{-3, 0, 2, 4, 7, 9},
		},
		{
			Input:    []int{2, 5, 1, 10, 12, 8},
			Expected: []int{1, 2, 5, 8, 10, 12},
		},
	}

	for _, caseData := range testData {
		testName := fmt.Sprintf("Sort %v ascending", caseData.Input)
		t.Run(testName, func(t *testing.T) {
			if res := QuickSortNaivePartition(caseData.Input); !compareArrays(res, caseData.Expected) {
				t.Errorf("Expected %v. Got %v", caseData.Expected, res)
			}
		})
	}
}

func TestQuickSortLomuto(t *testing.T) {
	testData := []testCase{
		{
			Input:    []int{9, 7, 4, 2, 0, -3},
			Expected: []int{-3, 0, 2, 4, 7, 9},
		},
		{
			Input:    []int{2, 5, 1, 10, 12, 8},
			Expected: []int{1, 2, 5, 8, 10, 12},
		},
	}

	for _, caseData := range testData {
		testName := fmt.Sprintf("Sort %v ascending", caseData.Input)
		t.Run(testName, func(t *testing.T) {
			if res := QuickSortLomutoPartition(caseData.Input); !compareArrays(res, caseData.Expected) {
				t.Errorf("Expected %v. Got %v", caseData.Expected, res)
			}
		})
	}
}

func TestQuickSortHoare(t *testing.T) {
	testData := []testCase{
		{
			Input:    []int{9, 7, 4, 2, 0, -3},
			Expected: []int{-3, 0, 2, 4, 7, 9},
		},
		{
			Input:    []int{2, 5, 1, 10, 12, 8},
			Expected: []int{1, 2, 5, 8, 10, 12},
		},
	}

	for _, caseData := range testData {
		testName := fmt.Sprintf("Sort %v ascending", caseData.Input)
		t.Run(testName, func(t *testing.T) {
			if res := QuickSortHoarePartition(caseData.Input); !compareArrays(res, caseData.Expected) {
				t.Errorf("Expected %v. Got %v", caseData.Expected, res)
			}
		})
	}
}
