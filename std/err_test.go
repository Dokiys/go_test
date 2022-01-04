package std

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

// TestErrWrap 包裹错误
func TestErrWrap(t *testing.T) {
	err := errors.New("This is some err")
	e1 := errors.Wrapf(err, "err2")
	t.Logf("e1: %s", e1)

	var nilErr error
	e2 := errors.Wrapf(nilErr, "err2")
	t.Logf("e2: %s", e2)			// e2也为nil
}

// TestErrTagError	// 标记错误
func TestErrTagError(t *testing.T) {
	e1 := errors.New("媒体账户不存在")
	e2 := errors.New("媒体账户不存在")
	t.Log(e1.Error() == e2.Error())
}

type MyErr struct { msg string }
func (self MyErr) Error() string     { return self.msg }
func (self MyErr) IsTemporary() bool { return true }

type Temporary interface { IsTemporary() bool }
// TestErrChain 错误链判断错误
func TestErrChain(t *testing.T) {
	err1 := MyErr{msg: "myErr1"}
	err2 := MyErr{msg: "myErr2"}
	var temp Temporary

	fmt.Printf("Is: %v\n", errors.Is(err1, err2))
	fmt.Printf("As: %v\n", errors.As(err1, &temp))
	fmt.Printf("As: %v\n", errors.As(err1, &MyErr{}))
}

func TestErrNil(t *testing.T) {
	var err error
	if _, ok := err.(*json.MarshalerError); ok {
		t.Log(1)
	}
}