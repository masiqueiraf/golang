package main

import "time"

func main() {
	canal1, canal2 := make(chan int), make(chan int)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- 1
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- 2
		}
	}()

	go func() {
		for {
			select {
			case msg := <-canal1:
				println(msg)
			case msg := <-canal2:
				println(msg)
			}
		}
	}()

	time.Sleep(time.Second * 5)
}
