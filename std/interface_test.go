package std

import "testing"

// TestInterfaceAssert 接口类型断言
func TestInterfaceAssert(t *testing.T) {
	var i interface{}
	str := "string"
	i = str

	if v, ok := i.(string); ok {
		t.Logf("i is string, value is: %s", v)
	}
}

type Aa struct {}
// TestInterfaceType 接口类型判断
func TestInterfaceType(t *testing.T) {
	var i interface{}
	//str := "string"
	a := Aa{}
	i = a

	switch i.(type) {
	case string:
		t.Logf("i is string, value is: %s", i.(string))
	case Aa:
		t.Logf("i is Aa, value is: %s", i.(Aa))
	default:
		t.Logf("unknow interface type")
	}
}

// TestInterfaceForcedTypeConv 强制类型转换
func TestInterfaceForcedTypeConv(t *testing.T) {
	var i int = 3
	var f float64

	f = float64(i)
	t.Logf("i forced conv to float: %f", f)
}