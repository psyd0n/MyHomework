package test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Test() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	var i int
	for {
		select {
		case <-sig:
			fmt.Printf("Завершаю работу, квадрат последнего числа %d", <-Duble(i))
			return
		default:
		}
		fmt.Print("Введите число или завершите работу:")
		fmt.Scan(&i)
		fmt.Println(<-Duble(i))
	}
}

func Duble(i int) chan int {
	ch := make(chan int)
	go func() {
		i *= i
		ch <- i
	}()
	return ch
}
