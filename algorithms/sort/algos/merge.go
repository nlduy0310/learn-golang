package algos

func MergeSort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	mergeSortRecursively(res, 0, len(res)-1)

	return res
}

func mergeSortRecursively(arr []int, left int, right int) {
	if !(left < right) {
		return
	}

	mid := left + (right-left)/2
	mergeSortRecursively(arr, left, mid)
	mergeSortRecursively(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	leftArr, rightArr := make([]int, mid-left+1), make([]int, right-mid)
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])

	cur := left
	i, j := 0, 0
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			arr[cur] = leftArr[i]
			i++
		} else {
			arr[cur] = rightArr[j]
			j++
		}
		cur++
	}

	for ; i < len(leftArr); i++ {
		arr[cur] = leftArr[i]
		cur++
	}

	for ; j < len(rightArr); j++ {
		arr[cur] = rightArr[j]
		cur++
	}
}
