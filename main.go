package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		//todo correct solution is to use while loop and increase the i++
		for i := 1; i < 5; i++ {
			time.Sleep(time.Millisecond)
			p := <-ch
			if p == 0 {
				fmt.Print(i*2, " ")
			}
			ch <- 1
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			p := <-ch
			if p == 1 {
				fmt.Print(i*2+1, " ")
			}
			ch <- 0

		}
	}()
	ch <- 1
	time.Sleep(time.Second)

}
