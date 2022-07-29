package mydoc_test

import (
	"fmt"

	"mydoc"
)

func ExampleVersion() {
	mydoc.Version()
	// Output:
	// v1.0.0
}

func ExampleStu_Study() {
	stu := mydoc.Stu{Name: "zhangsan"}
	fmt.Println(stu.Code())
	stu.Study()
	fmt.Println(stu.Code())
	// Output:
	// 0
	// 1
}

func ExampleMeet() {
	stu := mydoc.Stu{Name: "zhangsan"}
	teacher := mydoc.Teacher{Name: "lisi"}

	mydoc.Meet(teacher, stu)
	// Output:
	// lisi: Hello zhangsan
	// zhangsan: Hello Mr.lisi
}
