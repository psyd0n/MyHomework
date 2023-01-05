package test

import (
	"fmt"
)

func Duble(i int, ch chan int) chan int {
	go func() {
		i *= i
		ch <- i
	}()
	return ch
}

func Vvod(i int, ch chan int) {
	fmt.Print("Введите число или завершите работу:")
	fmt.Scan(&i)
	fmt.Println(<-Duble(i, ch))
}
