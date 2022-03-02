package main

import (
	"reflect"
	"testing"
)

// DistinctBy 根据传入方法比较去重
func DistinctBy(s interface{}, f interface{}) interface{} {
	return distinct(s, f)
}

// Distinct 根据基本类型相等比较去重
func Distinct(s interface{}) interface{} {
	return distinct(s, func(v interface{}) interface{} { return v })
}

func distinct(s interface{}, f interface{}) interface{} {
	sv := reflect.ValueOf(s)
	l := sv.Len()

	// 利用map去重
	m := make(map[interface{}]struct{}, l>>1)
	fv := reflect.ValueOf(f)
	iSlice:= make([]reflect.Value, 0, l>>1)
	for i := 0; i < l; i++ {
		v := sv.Index(i)
		k := fv.Call([]reflect.Value{v})[0].Interface()
		if _, ok := m[k]; ok {
			continue
		}
		m[k] = struct{}{}
		iSlice = append(iSlice, v)
	}
	res := reflect.MakeSlice(reflect.TypeOf(s), len(iSlice), len(iSlice))
	res = reflect.Append(res, iSlice...)

	return res.Interface()
}

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
