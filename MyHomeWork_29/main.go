package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type student struct {
	name  string
	age   int
	grade int
}

func (p student) info() string {
	return p.name + " " + strconv.Itoa(p.age) + " " + strconv.Itoa(p.grade)
}

type university struct {
	studentBy map[string]*student
}

func (u university) put(value *student) (int, error) {
	if value == nil {
		return -1, errors.New("student's value is nil")
	}

	if _, exist := u.studentBy[value.name]; !exist {
		u.studentBy[value.name] = value
		return 1, nil
	}

	return 0, nil
}

func (u university) get(name string) (*student, error) {
	if value, exist := u.studentBy[name]; !exist {
		return nil, fmt.Errorf("student %q is not exists", name)
	} else {
		return value, nil
	}
}

func main() {
	u := university{
		make(map[string]*student),
	}

	// os.Stdin, _ = os.Open("input.txt") // for testing purposes only

	var in = bufio.NewReader(os.Stdin)

	for {
		line, err := in.ReadString('\n')
		if err == io.EOF {
			fmt.Println("------------")
			fmt.Println("Студенты из хранилища:")
			break
		}

		var name string
		var age, grade int
		lineFields := strings.Fields(line)

		if len(lineFields) < 3 {
			fmt.Print("Некорректный ввод аргументов\n")
			continue
		}

		name = lineFields[0]
		age, errAge := strconv.Atoi(lineFields[1])
		grade, errGrade := strconv.Atoi(lineFields[2])

		if errAge != nil || errGrade != nil {
			fmt.Print("Ошибка")
			continue
		}

		student := &student{
			name:  name,
			age:   age,
			grade: grade,
		}

		if _, err := u.get(student.name); err != nil {
			u.put(student)
		} else {
			fmt.Print("Некорректный ввод данных\n")
		}
	}

	counter := 1
	for _, value := range u.studentBy {
		fmt.Println(counter, value.info())
		counter++
	}
}
