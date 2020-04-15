// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package algorithms

// ComparatorRune is a function that given two elements of type rune returns x such that:
// x < 0 iff e1 < e2
// x > 0 iff e1 > e2
// x = 0 iff e1 = e2
type ComparatorRune func(e1, e2 rune) int

func MergeSortRune(array []rune, compare ComparatorRune) {
	b := append([]rune(nil), array...) // shallow copy
	mergeSortRune(array, compare, 0, len(array)-1, b)
}

func mergeSortRune(array []rune, compare ComparatorRune, low, high int, b []rune) {
	if low >= high {
		return
	}

	middle := low + (high-low)/2
	mergeSortRune(array, compare, low, middle, b)
	mergeSortRune(array, compare, middle+1, high, b)
	mergeRune(array, compare, low, middle, high, b)
}

func mergeRune(array []rune, compare ComparatorRune, low int, middle int, high int, b []rune) {

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
