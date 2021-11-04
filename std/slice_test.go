package std

import (
	"sort"
	"testing"
)

// TestSliceRange range两个参数遍历
func TestSliceRange(t *testing.T) {
	s := []string{"a", "b", "c"}

	for i, v := range s { // index, value
		t.Log(i)
		t.Log(v)
	}

	for i := range s { // index
		t.Log(i) // 0, 1, 2
	}
}

// TestSliceRangePointer 遍历切片时，v 是同一个指针
func TestSliceRangePointer(t *testing.T) {
	type A struct{ Id int }
	slice := []A{{Id: 1}, {Id: 2}}

	for _, v := range slice {
		t.Logf("%p", &v)
	}
}

// TestSliceRangePointer2 遍历切片时，v 是同一个指针
func TestSliceRangePointer2(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	var index, value int
	for index, value = range slice {
		myMap[index] = &value
	}
	t.Log(myMap)
}

// TestSliceExpend 切片扩容
func TestSliceExpend(t *testing.T) {
	nums := make([]int32, 0, 2)

	t.Logf("len: %d",len(nums))
	t.Logf("cap: %d",cap(nums))
	t.Logf("--------")

	nums = append(nums, 1)
	t.Logf("len: %d",len(nums))
	t.Logf("cap: %d",cap(nums))
	t.Logf("--------")

	nums = append(nums, 2)
	t.Logf("len: %d",len(nums))
	t.Logf("cap: %d",cap(nums))
	t.Logf("--------")

	nums = append(nums, 3)
	t.Logf("len: %d",len(nums))
	t.Logf("cap: %d",cap(nums))
	t.Logf("--------")
}

// TestSliceSort 切片排序
func TestSliceSort(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3, 2, 1}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	t.Log(nums) // [1 1 2 2 3 3 4]
}
