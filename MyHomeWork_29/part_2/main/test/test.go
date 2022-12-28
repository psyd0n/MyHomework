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
	ch := make(chan int)
	for {
		select {
		case <-sig:
			fmt.Printf("Завершаю работу, квадрат последнего числа %d", <-Duble(i, ch))
			return
		default:
		}
		fmt.Print("Введите число или завершите работу:")
		fmt.Scan(&i)
		fmt.Println(<-Duble(i, ch))
	}
}

func Duble(i int, ch chan int) chan int {
	go func() {
		i *= i
		ch <- i
	}()
	return ch
}
