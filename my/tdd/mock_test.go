package tdd

import (
	"github.com/golang/mock/gomock"
	"testing"
)

// mockgen -source=./mytest/mock_test.go -destination=./mytest/mock_foo.go -package=mytest
type Foo interface {
	Bar(x int) int
}

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockFoo(ctrl)
	m.EXPECT().Bar(gomock.Eq(88)).Return(11)

	t.Logf("Expect arg return: %d",m.Bar(88))
	//t.Logf("Unexpect arg will raise error: %d",m.Bar(11))
}