package algos

func SelectionSort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		selected := i
		for j := i + 1; j < n; j++ {
			if res[j] < res[selected] {
				selected = j
			}
		}
		res[i], res[selected] = res[selected], res[i]
	}
	return res
}
