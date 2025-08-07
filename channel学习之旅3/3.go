package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(ch chan<- int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("生产者%d发送\n", i+1)
		ch <- i + 1
		time.Sleep(time.Millisecond * 500)
	}
	// 关闭管道
	fmt.Println("关闭ch")
	close(ch)
}

// 消费者
func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println("消费者接收到值")
		fmt.Println("<-ch", v)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
