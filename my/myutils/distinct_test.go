package myutils

import (
	"testing"
)

// TestDistinctInt 去重int
func TestDistinctInt(t *testing.T) {
	nums := []int{1, 2, 3, 4, 1, 2, 3, 1, 1, 1}

	res := Distinct(nums)

	t.Log(res)
}

// TestDistinctString 去重string
func TestDistinctString(t *testing.T) {
	nums := []string{"1", "2", "3", "4", "1", "2", "3", "1", "1", "1"}

	res := Distinct(nums)

	t.Log(res)
}

// TestDistinctByStruct 测试根据指定函数去重结构体
func TestDistinctByStruct(t *testing.T) {
	type a struct { Id int32 }
	as := []*a{
		{Id: 1},
		{Id: 1},
		{Id: 3},
	}

	res := DistinctBy(as, func(v *a) int32 {
		return v.Id
	}).([]*a)

	t.Log(res[0].Id)
}
