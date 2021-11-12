package main

import "time"

func demo(t1 interface{}, t2 interface{}) {
	for {
		select {
		case <-t1.(*time.Ticker).C:
			println("1s timer")

		case <-t2.(*time.Ticker).C:
			println("2s timer")
		}
	}
}

func main() {

	t1 := time.NewTicker(time.Second * 1)
	t2 := time.NewTicker(time.Second * 2)
	go demo(t1, t2)
	select {}
}