package myopt

import "testing"

type Student2 struct {
	Id    int
	Name  string
	Email string
}

type StudentBuilder2 struct {
	student *Student2
	opts     []Opt
}

type Opt func(*StudentBuilder2)

func Init2() *StudentBuilder2 {
	return &StudentBuilder2{
		student: &Student2{
			Id:    0,
			Name:  "",
			Email: "",
		},
		opts:    []Opt{},
	}
}

func (self *StudentBuilder2) Id(id int) *StudentBuilder2 {
	self.opts = append(self.opts, func(builder *StudentBuilder2) {
		builder.student.Id = id
	})
	return self
}

func (self *StudentBuilder2) Name(name string) *StudentBuilder2 {
	self.opts = append(self.opts, func(builder *StudentBuilder2) {
		builder.student.Name = name
	})
	return self
}

func (self *StudentBuilder2) Email(email string) *StudentBuilder2 {
	self.opts = append(self.opts, func(builder *StudentBuilder2) {
		builder.student.Email = email
	})
	return self
}

func (self *StudentBuilder2) Build() Student2 {
	for _, opt := range self.opts {
		opt(self)
	}
	return *self.student
}

func TestBuilder2(t *testing.T) {
	stu := Init2().Id(1).Name("zhangsan").Email("zhangsan@4399.com").Build()
	t.Logf("Student: %v", stu)
}
