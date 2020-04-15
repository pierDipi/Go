package algorithms

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMergeSort(t *testing.T) {

	tt := []struct {
		name     string
		array    []T
		expected []T
	}{
		{
			name:     "sorted",
			array:    []T{1, 2, 3, 4, 5},
			expected: []T{1, 2, 3, 4, 5},
		},
		{
			name:     "sorted ASC",
			array:    []T{1, 2, 3, 4, 5},
			expected: []T{1, 2, 3, 4, 5},
		},
		{
			name:     "sorted DESC",
			array:    []T{5, 4, 3, 2, 1},
			expected: []T{1, 2, 3, 4, 5},
		},
		{
			name:     "ASC_DESC_ASC_DESC",
			array:    []T{5, 10, 11, 3, 2, 5, 7, 0, -1},
			expected: []T{-1, 0, 2, 3, 5, 5, 7, 10, 11},
		},
		{
			name:     "DESC_ASC_DESC_ASC",
			array:    []T{10, 8, 6, 4, 3, 4, 5, 6, 7, 10, 8, 6, 4, 3, 4, 5, 6, 7, 10},
			expected: []T{3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 6, 7, 7, 8, 8, 10, 10, 10},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			MergeSortT(tc.array, func(e1, e2 T) int {
				return e1.(int) - e2.(int)
			})
			if diff := cmp.Diff(tc.array, tc.expected); diff != "" {
				t.Errorf("expected %v - got %v - diff %q", tc.expected, tc.array, diff)
			}
		})
	}

}
