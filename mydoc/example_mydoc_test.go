package mydoc_test

import (
	"fmt"
	"go_test/mydoc"
)

func Example() {
	mydoc.Version()
}

func ExampleStu_Study() {
	stu := mydoc.Stu{ Name: "zhangsan" }
	fmt.Println(stu.Code())
	stu.Study()
	fmt.Println(stu.Code())
}

func ExampleMeet() {
	stu := mydoc.Stu{ Name: "zhangsan" }
	teacher := mydoc.Teacher{ Name: "lisi" }

	mydoc.Meet(teacher, stu)
}
