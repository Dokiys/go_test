package std

import "testing"

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
