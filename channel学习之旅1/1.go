package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	go func() {
		fmt.Println("[Goroutine] Send 1")
		c <- 1
		fmt.Println("[Goroutine] Send 2")
		c <- 2
		fmt.Println("[Goroutine] Send 2")
		c <- 3
	}()

	time.Sleep(time.Millisecond * 10) // 稍等，确保 goroutine 启动

	fmt.Println("[Main] Receive 1")
	v1 := <-c
	fmt.Println("[Main] Got:", v1)

	fmt.Println("[Main] Receive 2")
	v2 := <-c
	fmt.Println("[Main] Got:", v2)
}
