package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"v/lesson/MyHomework/MyHomework_29/part_2/main/test"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	var i int
	ch := make(chan int)
	for {
		select {
		case <-sig:
			fmt.Printf("Завершаю работу, квадрат последнего числа %d", <-test.Duble(i, ch))
			return
		default:
		}
		test.Vvod(i, ch)
	}
}
