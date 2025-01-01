package algos

func HeapSort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	n := len(res)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(res, n, i)
	}

	for i := 0; i < n-1; i++ {
		res[0], res[n-1-i] = res[n-1-i], res[0]
		heapify(res, n-1-i, 0)
	}

	return res
}

func heapify(arr []int, n, i int) {
	maxIdx := i

	leftIdx, rightIdx := 2*i+1, 2*i+2

	if leftIdx < n && arr[leftIdx] > arr[maxIdx] {
		maxIdx = leftIdx
	}

	if rightIdx < n && arr[rightIdx] > arr[maxIdx] {
		maxIdx = rightIdx
	}

	if maxIdx != i {
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
		heapify(arr, n, maxIdx)
	}
}
