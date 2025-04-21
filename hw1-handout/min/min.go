package min

// Min returns the minimum value in the arr,
// and 0 if arr is nil.
func Min(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
