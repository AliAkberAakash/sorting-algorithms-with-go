package quicksort_test

import (
	"testing"

	"github.com/AliAkberAakash/sorting-algorithms-with-go/quicksort"
)

func TestQuickSortUnsorted(t *testing.T) {
	unsorted := []int{
		10, 7, 2, 5, 0, 8, 6, 9, 1, 3, 4,
	}

	exp := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	quicksort.QuickSort(&unsorted, 0, len(unsorted)-1)

	for i := 0; i < 11; i++ {
		if unsorted[i] != exp[i] {
			t.Errorf("Expected %d found %d\n", unsorted[i], exp[i])
		}
	}

}

func TestQuickSortSorted(t *testing.T) {
	sorted := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	exp := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	quicksort.QuickSort(&sorted, 0, len(sorted)-1)

	for i := 0; i < 11; i++ {
		if sorted[i] != exp[i] {
			t.Errorf("Expected %d found %d\n", sorted[i], exp[i])
		}
	}

}
