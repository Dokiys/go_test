package mydoc

import "fmt"

// version mydoc version
const version = "v1.0"

/*
	The following is package method
	This annotation will not appear in go doc, because of blank lines
*/

// Meet greeting each other
func Meet(teacher Teacher, stu Stu)  {
	fmt.Println(teacher.Name + ": Hello " + stu.Name)
	fmt.Println(stu.Name + ": Hello Mr." + teacher.Name )
}

// Version show mydoc version
func Version()  {
	fmt.Println(Version)
}

// Stu a student struct
type Stu struct {
	Name string
	code int
}

// Study add Code for Stu
func (self *Stu)Study() {
	self.code += 1
}

// Code return stu code
func (self *Stu)Code() int {
	return self.code
}

// Teacher a Teacher struct
type Teacher struct {
	Name string
}

// BUG(Dokiy): Unused
func (self *Teacher)evaluate(stu *Stu, code int) {
	stu.code = code
}