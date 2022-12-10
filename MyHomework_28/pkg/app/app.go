package app

import (
	st "Myhomework_28/pkg/storage"
	"Myhomework_28/pkg/student"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Run() {
	var storage st.Storage
	storage = st.NewUniversity()
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
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

		if _, err := storage.Get(student.Name); err != nil {
			storage.Put(student)
		} else {
			fmt.Print("Некорректный ввод данных\n")
		}
	}

	counter := 1
	for _, value := range storage.GetStudents() {
		fmt.Println(counter, value.Info())
		counter++
	}
}
