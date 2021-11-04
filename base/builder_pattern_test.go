package base

import "testing"

type Student struct {
	Id    int
	Name  string
	Email string
}

type StudentBuilder struct {
	student *Student
}

func Init() *StudentBuilder {
	return &StudentBuilder{
		student: &Student{
			Id:    0,
			Name:  "",
			Email: "",
		},
	}
}

func (self *StudentBuilder)Id(id int) *StudentBuilder {
	self.student.Id = id
	return self
}

func (self *StudentBuilder)Name(name string) *StudentBuilder {
	self.student.Name = name
	return self
}

func (self *StudentBuilder)Email(email string) *StudentBuilder {
	self.student.Email = email
	return self
}

func (self *StudentBuilder)Build() Student {
	return *self.student
}

func TestBuilder(t *testing.T) {
	stu := Init().Id(1).Name("zhangsan").Email("zhangsan@4399.com").Build()
	t.Logf("Student: %v", stu)
}