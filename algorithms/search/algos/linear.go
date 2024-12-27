package algos

func LinearSearch(array []int, value int) (bool, int) {
	n := len(array)

	for i := 0; i < n; i++ {
		if array[i] == value {
			return true, i
		}
	}

	return false, -1
}
