package algos

import (
	"math/rand"
)

func QuickSortNaivePartition(input []int) []int {
	res := make([]int, len(input))
	copy(res, input)

	partition := func(arr []int, left, right, pivotIndex int) int {
		temp := make([]int, right-left+1)
		copy(temp, arr[left:right+1])

		pivotValue := arr[pivotIndex]
		x, X := left, right
		for i := 0; i < len(temp); i++ {
			if i == pivotIndex-left {
				continue
			}

			if temp[i] < pivotValue {
				arr[x] = temp[i]
				x++
			} else {
				arr[X] = temp[i]
				X--
			}
		}
		arr[x] = pivotValue
		return x
	}
	// sort in-place and recursively
	var quickSortRecursively func(arr []int, left, right int)
	quickSortRecursively = func(arr []int, left, right int) {
		if !(left < right) {
			return
		}

		pivotIndex := selectPivot(left, right, Random)
		newPivotIndex := partition(arr, left, right, pivotIndex)
		quickSortRecursively(arr, left, newPivotIndex-1)
		quickSortRecursively(arr, newPivotIndex+1, right)
	}

	quickSortRecursively(res, 0, len(res)-1)

	return res
}

func QuickSortLomutoPartition(input []int) []int {
	res := make([]int, len(input))
	copy(res, input)

	partition := func(arr []int, left, right int) int {
		i, j := left-1, left
		pivotValue := arr[right]
		for ; j < right; j++ {
			if arr[j] >= pivotValue {
				continue
			} else {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		i++
		arr[i], arr[right] = arr[right], arr[i]
		return i
	}
	// sort in-place and recursively
	var quickSortRecursively func(arr []int, left, right int)
	quickSortRecursively = func(arr []int, left, right int) {
		if !(left < right) {
			return
		}

		newPivotIndex := partition(arr, left, right)
		quickSortRecursively(arr, left, newPivotIndex-1)
		quickSortRecursively(arr, newPivotIndex+1, right)
	}

	quickSortRecursively(res, 0, len(res)-1)

	return res
}

func QuickSortHoarePartition(input []int) []int {
	res := make([]int, len(input))
	copy(res, input)

	partition := func(arr []int, left, right int) int {
		pivotValue := arr[left]
		i, j := left, right
		for {
			for ; i < right; i++ {
				if arr[i] >= pivotValue {
					break
				}
			}

			for ; j > left; j-- {
				if arr[j] < pivotValue {
					break
				}
			}

			if i >= j {
				return j
			}

			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// sort in-place and recursively
	var quickSortRecursively func(arr []int, left, right int)
	quickSortRecursively = func(arr []int, left, right int) {
		if !(left < right) {
			return
		}

		newPivotIndex := partition(arr, left, right)
		quickSortRecursively(arr, left, newPivotIndex)
		quickSortRecursively(arr, newPivotIndex+1, right)
	}

	quickSortRecursively(res, 0, len(res)-1)

	return res
}

type PivotSelectionMode int

const (
	LeftMost PivotSelectionMode = iota
	RightMost
	Middle
	Random
)

// random pivot
func selectPivot(left, right int, selectionMode PivotSelectionMode) int {
	switch selectionMode {
	case LeftMost:
		return left
	case RightMost:
		return right
	case Middle:
		return left + (right-left)/2
	default:
		return rand.Intn(right-left+1) + left
	}
}
