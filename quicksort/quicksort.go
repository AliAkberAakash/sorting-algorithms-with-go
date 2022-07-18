package quicksort

func partition(arr *[]int, start, end int) int {
	pivot := (*arr)[end]

	largest := start - 1

	for start < end {
		if (*arr)[start] <= pivot {
			largest++
			(*arr)[start], (*arr)[largest] = (*arr)[largest], (*arr)[start]
		}
		start++
	}

	(*arr)[largest+1], (*arr)[end] = (*arr)[end], (*arr)[largest+1]

	return largest + 1
}

func QuickSort(arr *[]int, start, end int) {
	if start < end {
		p := partition(arr, start, end)

		QuickSort(arr, p+1, end)
		QuickSort(arr, start, p-1)
	}

}
