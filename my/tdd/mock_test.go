package tdd

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -source=./mock_test.go -destination=./mock_inf.go -package=tdd
type Inf interface {
	Bar() int
	IsGood(name string) bool
	IsOldPerson(p *People) bool
}

type People struct {
	age int
}

type B struct{}

func (self *B) Bar() int {
	return rand.Intn(2)
}

type A struct {
	inf Inf
}

func (self *A) IsOne() bool {
	if r := self.inf.Bar(); r == 1 {
		return true
	} else {
		return false
	}
}

type InfTest1 struct{}

func (self *InfTest1) Bar() int {
	return 1
}

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInf := NewMockInf(ctrl)
	mockInf.EXPECT().Bar().Return(1).AnyTimes()
	mockInf.EXPECT().IsGood(gomock.Eq("zhangsan")).Return(true).AnyTimes()
	mockInf.EXPECT().IsOldPerson(newGtAgeMatcher(10)).Return(true).AnyTimes()
	a := &A{inf: mockInf}

	assert.Equal(t, true, a.IsOne())
}

type gtAgeMatcher struct {
	age int
}

func newGtAgeMatcher(age int) *gtAgeMatcher {
	return &gtAgeMatcher{age: age}
}

func (self *gtAgeMatcher) Matches(x interface{}) bool {
	p, ok := x.(*People)
	if !ok {
		return false
	}
	if self.age < p.age {
		return true
	}

	return false
}

func (self *gtAgeMatcher) String() string {
	return fmt.Sprintf("age > %d", self.age)
}
