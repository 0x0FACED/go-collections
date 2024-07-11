package list

const (
	TimSort    = 0
	QuickSort  = 1
	MergeSort  = 2
	BubbleSort = 3
)

const minMerge = 32

func minRunLength(n int) int {
	r := 0
	for n >= minMerge {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func (a *arrayList[T]) insertionSort(left, right int, compare Comparator[T]) {
	for i := left + 1; i <= right; i++ {
		temp := a.items[i]
		j := i - 1
		for j >= left && compare(temp, a.items[j]) {
			a.items[j+1] = a.items[j]
			j--
		}
		a.items[j+1] = temp
	}
}

func (a *arrayList[T]) merge(l, m, r int, compare Comparator[T]) {
	len1, len2 := m-l+1, r-m
	left := make([]T, len1)
	right := make([]T, len2)

	copy(left, a.items[l:l+len1])
	copy(right, a.items[m+1:m+1+len2])

	i, j, k := 0, 0, l

	for i < len1 && j < len2 {
		if compare(left[i], right[j]) {
			a.items[k] = left[i]
			i++
		} else {
			a.items[k] = right[j]
			j++
		}
		k++
	}

	for i < len1 {
		a.items[k] = left[i]
		i++
		k++
	}

	for j < len2 {
		a.items[k] = right[j]
		j++
		k++
	}
}

func (a *arrayList[T]) timSort(compare Comparator[T]) {
	n := a.size
	minRun := minRunLength(minMerge)

	for start := 0; start < n; start += minRun {
		end := start + minRun - 1
		if end >= n {
			end = n - 1
		}
		a.insertionSort(start, end, compare)
	}

	for size := minRun; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := left + size - 1
			right := min(left+2*size-1, n-1)

			if mid < right {
				a.merge(left, mid, right, compare)
			}
		}
	}
}

func (a *arrayList[T]) mergeSort(compare Comparator[T]) {
	a.mergeSortHelper(0, a.size-1, compare)
}

func (a *arrayList[T]) mergeSortHelper(l, r int, compare Comparator[T]) {
	if l < r {
		m := l + (r-l)/2

		a.mergeSortHelper(l, m, compare)
		a.mergeSortHelper(m+1, r, compare)

		a.merge(l, m, r, compare)
	}
}

func (a *arrayList[T]) quickSort(compare Comparator[T]) {
	a.quickSortHelper(0, a.size-1, compare)
}

func (a *arrayList[T]) quickSortHelper(l, r int, compare Comparator[T]) {
	if l < r {
		p := a.partition(l, r, compare)
		a.quickSortHelper(l, p-1, compare)
		a.quickSortHelper(p+1, r, compare)
	}
}

func (a *arrayList[T]) partition(l, r int, compare Comparator[T]) int {
	pivot := a.items[r]
	i := l - 1

	for j := l; j < r; j++ {
		if compare(a.items[j], pivot) {
			i++
			a.items[i], a.items[j] = a.items[j], a.items[i]
		}
	}
	a.items[i+1], a.items[r] = a.items[r], a.items[i+1]
	return i + 1
}

func (a *arrayList[T]) bubbleSort(compare Comparator[T]) {
	n := a.size
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if compare(a.items[j+1], a.items[j]) {
				a.items[j], a.items[j+1] = a.items[j+1], a.items[j]
			}
		}
	}
}
