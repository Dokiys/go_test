package std

import (
	"fmt"
	"testing"
)

// TestPointer 指向指针的指针
func TestPointer(t *testing.T) {
	type A struct{}

	a := &A{}
	b := &a
	t.Logf("type b:%T", b)

	c := &b
	t.Logf("type c:%T", c)

	t.Log("-----")
	t.Logf("pointer a:%p", a)
	t.Logf("pointer b:%p", b)
	t.Logf("pointer c:%p", c)
}

type A struct {
	b *B
}
type B struct {
	i int32
}

// TestPointerPassByPointer 测试指针传递
func TestPointerPassByPointer(t *testing.T) {
	b := &B{i:1}
	a := &A{b:b}
	fmt.Printf("before pass by pointer \n")
	fmt.Printf("\tfunc a is %p\n",&a)
	fmt.Printf("\tfunc a.b is %v\n",a.b)
	passByPointer(a)
	fmt.Printf("after pass by pointer \n")
	fmt.Printf("\tfunc a is %p\n",&a)
	fmt.Printf("\tfunc a.b is %v\n",a.b)
}

// TestPointerPassByValue 测试值传递
func TestPointerPassByValue(t *testing.T) {
	b := &B{i:1}
	a := &A{b:b}
	fmt.Printf("before pass by value \n")
	fmt.Printf("\tfunc a is %p\n",&a)
	fmt.Printf("\tfunc a.b is %v\n",a.b)
	passByValue(*a)
	fmt.Printf("after pass by value \n")
	fmt.Printf("\tfunc a is %p\n",&a)
	fmt.Printf("\tfunc a.b is %v\n",a.b)
}

func passByPointer(a *A) {
	b := B{i:2}
	a.b = &b
}

func passByValue(a A) {
	b := B{i:2}
	a.b = &b
}