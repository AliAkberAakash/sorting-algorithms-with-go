package mergesort

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var tmp []int

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}

	if i < len(left) {
		tmp = append(tmp, left[i:]...)
	}

	if j < len(right) {
		tmp = append(tmp, right[j:]...)
	}

	return tmp
}
