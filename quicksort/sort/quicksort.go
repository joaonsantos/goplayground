package sort

// partition finds correct pivot position in slice s with bounds l(ow) and h(igh)
// pivot is always the last element of the slice
//
// the idea is to use a "wall" which contains all elements smaller than the pivot
// an item is swapped when it is smaller than the pivot because we want it to be behind the wall
// when an element is swapped, advance the wall to fit more space
func partition(s []int, l, h int) int {
	pv := s[h]
	wall := l

	for i := l + 1; i < h; i++ {
		if s[i] <= pv {
			s[wall], s[i] = s[i], s[wall]
			wall++
		}
	}

	// useful if the array turns out to be already sorted
	// the pivot should be of a higher value than the wall value before swapping
	if s[h] < s[wall] {
		s[wall], s[h] = s[h], s[wall]
	}

	return wall
}

// function quicksort sorts slice s with bounds l(ow) and h(igh)
func quicksort(s []int, l, h int) {
	if l < h {
		p := partition(s, l, h)
		quicksort(s, l, p)
		quicksort(s, p+1, h)
	}
}

// function Quicksort is sorts a slice s and returns slice s(orted)s(slice)
func Quicksort(s []int) []int {
	ss := append([]int{}, s...)
	quicksort(ss, 0, len(ss)-1)

	return ss
}
