package mergesort_test

import (
	"testing"

	"github.com/AliAkberAakash/sorting-algorithms-with-go/mergesort"
)

func TestMergeSort(t *testing.T) {
	var array = []int{
		2, 9, 3, 1, 7, 0, 10, 4, 8, 5, 6,
	}

	var expected = []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	array = mergesort.MergeSort(array)

	for i := 0; i < len(array); i++ {
		if array[i] != expected[i] {
			t.Errorf("Expected %d found %d\n", expected[i], array[i])
		}
	}

}
