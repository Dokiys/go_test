package tdd

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestConvey(t *testing.T) {
	Convey("3 add 6 should equal 9", t, func() {
		sum := add(3, 6)
		So(sum, ShouldEqual, 11)
	})

	Convey("3 add 2 should equal 5", t, func() {
		sum := add(3, 2)
		So(sum, ShouldEqual, 6)
	})

	Convey("1 add 1 should equal 2", t, func() {
		sum := add(1, 1)
		So(sum, ShouldEqual, "string")
	})
}
