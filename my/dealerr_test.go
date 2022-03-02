package main

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type MyErr struct {	msg string }
func (self MyErr) Error() string     { return self.msg }
func (self MyErr) IsTemporary() bool { return true }

type Temporary interface { IsTemporary() bool }

func IsTemporary(err error) bool {
	// v1
	//te, ok := err.(Temporary)
	//return ok && te.IsTemporary()

	// v2
	var t Temporary
	return errors.As(err, &t)
}

func doSomething() error { return MyErr{msg: "my err"} }

// TestRaiseErr 直接抛出错误
func TestRaiseErr(t *testing.T) {
	err := doSomething()
	if err != nil {
		//return fmt.Errorf("call doSomething() Err: %s", err.Error())
		e :=  errors.New(fmt.Sprintf("call doSomething() Err: %s\n\t", err.Error()))
		t.Log(e)
	}
}

// TestWrapErr 包裹错误再抛出
func TestWrapErr(t *testing.T) {
	err := doSomething()
	if err != nil {
		e := errors.Wrapf(err, "call doSomething() Err:\n\t")
		t.Log(e)
	}
}

// TestDealTheErr 处理指定类型的错误
func TestDealTheErr(t *testing.T) {
	err := errors.Wrapf(doSomething(), "call doSomething() Err:\n\t")
	if err == nil {
		return
	}
	myErr, ok := errors.Cause(err).(MyErr)
	if ok {
		t.Logf("MyErr: %v", myErr)
	} else {
		t.Logf("Error: %v", myErr)
	}
}

// TestDealTheBehavior 根据错误行为来判断处理错误
func TestDealTheBehavior(t *testing.T) {
	err := errors.Wrapf(doSomething(), "call doSomething() Err:\n\t")
	if IsTemporary(err) {
		t.Logf("TemporaryError: %v\n\t", err)
	} else {
		t.Logf("Error: %v\n\t", err)
	}
}
