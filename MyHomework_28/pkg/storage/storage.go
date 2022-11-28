package storage

import (
	"Myhomework_28/pkg/student"
	"errors"
	"fmt"
)

type University struct {
	StudentBy map[string]*student.Student
}

func NewUniversity(StudentBy map[string]*student.Student) *University {
	return &University{
		StudentBy: StudentBy,
	}
}

func (u University) Put(value *student.Student) (int, error) {
	if value == nil {
		return -1, errors.New("student's value is nil")
	}

	if _, exist := u.StudentBy[value.Name]; !exist {
		u.StudentBy[value.Name] = value
		return 1, nil
	}

	return 0, nil
}

func (u University) Get(name string) (*student.Student, error) {
	if value, exist := u.StudentBy[name]; !exist {
		return nil, fmt.Errorf("student %q is not exists", name)
	} else {
		return value, nil
	}
}
