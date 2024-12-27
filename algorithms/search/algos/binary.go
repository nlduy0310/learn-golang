package algos

// Assuming that the array is sorted ascending
func BinarySearch(array []int, value int) (bool, int) {
	var (
		found = false
		index = -1
	)
	lower, upper := 0, len(array)-1

searchLoop:
	for lower <= upper {
		mid := (lower + upper) / 2
		switch {
		case array[mid] > value:
			upper = mid - 1
		case array[mid] < value:
			lower = mid + 1
		default:
			found = true
			index = mid
			break searchLoop
		}
	}

	return found, index
}
