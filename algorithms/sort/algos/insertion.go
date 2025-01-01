package algos

func InsertionSort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for i := 1; i < len(res); i++ {
		insert(res, i)
	}

	return res
}

func insert(arr []int, idx int) {
	insertAt := 0
	for i := 0; i < idx; i++ {
		if arr[i] >= arr[idx] {
			insertAt = i
			break
		}
	}

	if !(arr[insertAt] >= arr[idx]) {
		return
	}

	tmp := arr[idx]
	for i := idx; i > insertAt; i-- {
		arr[i] = arr[i-1]
	}
	arr[insertAt] = tmp
}
