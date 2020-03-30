package sort

func bubblesort(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out)-i-1; j++ {
			if out[j] > out[j+1] {
				out[j], out[j+1] = out[j+1], out[j]
			}
		}
	}
	return out
}

func mergesort(in []int) []int {
	values := make([]int, len(in))
	copy(values, in)
	if len(values) == 1 {
		return values
	}
	pivot := len(values) / 2
	return merge(
		mergesort(values[:pivot]),
		mergesort(values[pivot:]),
	)
}

func merge(left, right []int) []int {
	size := len(left) + len(right)
	values := make([]int, size, size)
	for i, j, k := 0, 0, 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			values[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			values[k] = left[i]
			i++
		} else if left[i] < right[j] {
			values[k] = left[i]
			i++
		} else {
			values[k] = right[j]
			j++
		}
	}
	return values
}
