// Package sort provides primitives for sorting slices.
package sort

// merge merges two slices left and right and returns a new sorted slice s
// the ideia is to create a temp slice to store the lowest number between the two slices
// and iterate them both, taking special care to merge the remaining elements
func merge(left []int, right []int) []int {
	i, leftEnd := 0, len(left)
	j, rightEnd := 0, len(right)
	curr, s := 0, make([]int, rightEnd+leftEnd)

	for i < leftEnd && j < rightEnd {
		if left[i] <= right[j] {
			s[curr] = left[i]
			i++

		} else {
			s[curr] = right[j]
			j++
		}
		curr++
	}

	// merge remaining elements
	for i < leftEnd {
		s[curr] = left[i]
		i++
		curr++
	}

	for j < rightEnd {
		s[curr] = right[j]
		j++
		curr++
	}

	return s

}

// Mergesort sorts a slice s and returns slice s(orted)s(slice).
// It returns a sorted slice.
func Mergesort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2
	return merge(Mergesort(s[:mid]), Mergesort(s[mid:]))
}
