package reflect

import (
	"reflect"
	"testing"
)

type Student struct {
	Id int
	Name string
	Age int
}

// TestReflect 测试反射
func TestReflect(t *testing.T) {
	stu := Student{9527, "Andy", 18}
	var i interface{}
	i = stu

	// 获取变量的 reflect.Type
	reType := reflect.TypeOf(i)
	t.Logf("reflect.Type=%s\n", reType)

	// 获取变量的 reflect.Value
	reVal := reflect.ValueOf(i)
	t.Logf("reflect.Value=%s\n", reVal)

	//打印reVal类型，使用 reVal，打印Name 成员 失败。无法索引Name成员
	//fmt.Printf("reVal=%T, name=%v",reVal,  reVal.Name)
	t.Logf("reVal=%T\n",reVal)

	// 将 reVal 转成 interface
	iVal := reVal.Interface()
	t.Logf("iVal= %v, type= %T\n", iVal, iVal)
	// iVal.Name 会报错Unresolved reference 'Name'
	// fmt.Printf("iVal= %v, type= %T, name= %v\n", iVal, iVal, iVal.Name)

	// 将 interface 通过类型断言 转回成 Student
	// stu:= iVal.(Student)
	if stu, ok := iVal.(Student); ok {
		t.Logf("stu= %v, type= %T, name=%v\n", stu, stu, stu.Name)
	}
}

// TestReflectCall 反射调用函数
func TestReflectCall(t *testing.T) {
	f := func(v int) {
		t.Logf("get value: %d",v)
	}
	num := 123

	var iNum interface{}
	var iF interface{}

	iF = f
	iNum = num

	valueF := reflect.ValueOf(iF)
	params := []reflect.Value{reflect.ValueOf(iNum)}

	result := valueF.Call(params)
	t.Logf("result len: %d",len(result))
}

// TestReflectSliceIndex 测试反射index切片
func TestReflectSliceIndex(t *testing.T) {
	f := func(v int) {
		t.Log(v)
	}

	valueF := reflect.ValueOf(f)

	nums := []int{1, 2, 3}
	var i interface{}
	i = nums

	valueI := reflect.ValueOf(i)
	for i := 0; i < valueI.Len(); i++ {
		valueF.Call([]reflect.Value{valueI.Index(i)})
	}
}

// TestReflectElem 反射Elem
func TestReflectElem(t *testing.T) {
	nums := []int{1, 2, 3}

	t.Log(reflect.TypeOf(nums).Elem().String())
}