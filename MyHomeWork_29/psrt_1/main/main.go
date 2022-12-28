package main

import (
	"fmt"
)

func main() {
	var str string
	var i int

	for {
		fmt.Print("введите число или \"стоп\": ")
		fmt.Scanln(&str)

		if str == "стоп" {
			break
		} else {
			if _, err := fmt.Sscan(str, &i); err == nil {
				fmt.Println("вы ввели ", i)
				zero := ZeroGo(i)
				one := OneGo(zero)
				two := TwoGo(one)
				fmt.Println(<-two)
			} else {
				fmt.Println("вы ввели не число")
			}
		}
	}
}

func ZeroGo(i int) chan int {
	firstChan := make(chan int)
	go func() {
		firstChan <- i
	}()
	return firstChan
}
func OneGo(ch chan int) chan int {
	secondChan := make(chan int)
	go func() {
		f := <-ch
		f *= f
		fmt.Println(f)
		secondChan <- f
	}()
	return secondChan
}

func TwoGo(cha chan int) chan int {
	ch := make(chan int)
	go func() {
		f := <-cha
		f *= 2
		ch <- f
	}()
	return ch
}
