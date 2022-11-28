package main

import (
	"Myhomework_28/pkg/storage"
	"Myhomework_28/pkg/student"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	mar := storage.NewUniversity(make(map[string]*student.Student))
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

		student := student.NewStudent(name, age, grade)

		if _, err := mar.Get(student.Name); err != nil {
			mar.Put(student)
		} else {
			fmt.Print("Некорректный ввод данных\n")
		}
	}

	counter := 1
	for _, value := range mar.StudentBy {
		fmt.Println(counter, value.Info())
		counter++
	}
}
