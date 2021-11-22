package mytest

import (
	"fmt"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestIt(t *testing.T) {
	It("should eq 2 while 1 add 1", t, func(a Assertions) {
		sum := add(1, 1)

		a.Expect(sum).Equal(3)
	})
}

func It(description string, t *testing.T, f func(assert Assertions)) {
	assert := &Assertions{
		t:   t,
		msg: description,
	}
	f(*assert)
}

type Assertions struct {
	t     *testing.T
	value interface{}
	msg   string
}

func (self *Assertions) Expect(v interface{}) *Assertions {
	self.value = v
	return self
}

// TODO[Dokiy] 2021/11/5: 利用调用栈，输出调用位置
func (self *Assertions) Equal(v interface{}) *Assertions {
	//self.value == v
	//ok, msg := equal(self.value, v)
	ok := (self.value.(int) == v.(int))
	if !ok {
		panic(fmt.Sprintf("Fatal %s ×\n\tActual(%d)", self.msg, self.value.(int)))
	}
	return self
}