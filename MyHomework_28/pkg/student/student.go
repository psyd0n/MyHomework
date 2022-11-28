package student

import "strconv"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(name string, age, grade int) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}

func (p Student) Info() string {
	return p.Name + " " + strconv.Itoa(p.Age) + " " + strconv.Itoa(p.Grade)
}
