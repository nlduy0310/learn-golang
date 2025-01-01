package algos

func BubbleSortNaive(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res)-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}

	return res
}

func BubbleSortCheckSwaps(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

outerLoop:
	for i := 0; i < len(res); i++ {
		swap := false
		for j := 0; j < len(res)-1; j++ {
			if res[j] > res[j+1] {
				swap = true
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
		if !swap {
			break outerLoop
		}
	}

	return res
}
