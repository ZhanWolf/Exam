package main

import "fmt"

func main() {

	var ch = make(chan int, 10)

	for i := 0; i < cap(ch); i++ {
		go func() {
			fmt.Println("有点强人锁男")
			ch <- 1
		}()
	}

	for i := 0; i < cap(ch); i++ {
		<-ch
	}
}
