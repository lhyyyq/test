package main

import (
	"fmt"
	"time"
)

var POOL = 100

func go1(p chan int) {

	for i := 1; i <= POOL; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("groutine-1:", i)
		}
	}
}

func go2(p chan int) {

	for i := 1; i <= POOL; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("groutine-2:", i)
		}
	}
}

func main() {
	msg := make(chan int)

	go go1(msg)
	go go2(msg)

	time.Sleep(time.Second * 1)

}
