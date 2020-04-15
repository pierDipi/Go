package algorithms

import "github.com/cheekybits/genny/generic"

type T generic.Type

// ComparatorT is a function that given two elements of type T returns x such that:
// x < 0 iff e1 < e2
// x > 0 iff e1 > e2
// x = 0 iff e1 = e2
type ComparatorT func(e1, e2 T) int

func MergeSortT(array []T, compare ComparatorT) {
	b := append([]T(nil), array...) // shallow copy
	mergeSortT(array, compare, 0, len(array)-1, b)
}

func mergeSortT(array []T, compare ComparatorT, low, high int, b []T) {
	if low >= high {
		return
	}

	middle := low + (high-low)/2
	mergeSortT(array, compare, low, middle, b)
	mergeSortT(array, compare, middle+1, high, b)
	mergeT(array, compare, low, middle, high, b)
}

func mergeT(array []T, compare ComparatorT, low int, middle int, high int, b []T) {

	for i := low; i <= middle; i++ {
		b[i] = array[i]
	}
	for i := middle + 1; i <= high; i++ {
		b[i] = array[i]
	}

	lowQueue := low
	highQueue := middle + 1
	i := low
	for ; lowQueue <= middle && highQueue <= high; i++ {
		if compare(b[lowQueue], b[highQueue]) > 0 { // b[lowQueue] > b[highQueue]
			array[i] = b[highQueue]
			highQueue++
		} else {
			array[i] = b[lowQueue]
			lowQueue++
		}
	}

	for ; lowQueue <= middle; i++ {
		array[i] = b[lowQueue]
		lowQueue++
	}

	for ; highQueue <= high; i++ {
		array[i] = b[highQueue]
		highQueue++
	}
}
