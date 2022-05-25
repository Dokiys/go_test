package alg

// 堆排序
func HeapSort(nums []int) {
	end := len(nums) - 1
	for i := end >> 1; i >= 0; i-- {
		sink(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[end] = nums[end], nums[0]
		end--
		sink(nums, 0, end)
	}

}

func sink(heap []int, root, end int) {
	for {
		child := (root << 1) + 1
		if child > end {
			return
		}
		if child < end && heap[child] < heap[child+1] {
			child++
		}
		if heap[root] >= heap[child] {
			return
		}
		heap[root], heap[child] = heap[child], heap[root]
		root = child
	}
}

// 分治排序
func MergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	mid := len(nums) >> 1
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	result := make([]int, len(nums))
	p1, p2 := 0, 0
	i := 0

	for p1 < len(left) && p2 < len(right) {
		if left[p1] <= right[p2] {
			result[i] = left[p1]
			p1++
		} else {
			result[i] = right[p2]
			p2++
		}
		i++
	}

	copy(result[i:], left[p1:])
	copy(result[i:], right[p2:])
	return result
}

// 快速排序
func QuickSort(s []int) {
	if len(s) <= 1 {
		return
	}
	head, tail := 1, len(s)-1
	for head < tail {
		if s[head] > s[0] {
			s[tail], s[head] = s[head], s[tail]
			tail--
		} else {
			head++
		}
	}
	if s[head] > s[0] {
		head--
	}

	s[head], s[0] = s[0], s[head]
	QuickSort(s[:head])
	QuickSort(s[head+1:])
}

// 找到一个随机数组中，第n大的数
func FindMaxn(s []int, n int) int {
	if n > len(s) || n <= 0 {
		panic("Out of index")
	}

	head, tail := 1, len(s)-1
	for head < tail {
		if s[head] > s[0] {
			s[tail], s[head] = s[head], s[tail]
			tail--
		} else {
			head++
		}
	}
	if s[head] > s[0] {
		head--
	}
	s[head], s[0] = s[0], s[head]

	if head+1 > n {
		return FindMaxn(s[:head+1], n)
	} else if head+1 == n {
		return s[head]
	} else {
		return FindMaxn(s[head+1:], n-head-1)
	}
}