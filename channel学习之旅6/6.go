package main

import (
	"fmt"
)

func Print10(sem chan<- struct{}) {
	for i := 0; i < 10; i++ {
		fmt.Println("开始打印", i)
		fmt.Println("通知信号量完成")
		sem <- struct{}{}
	}
}

func main() {
	sem := make(chan struct{}, 10)
	go Print10(sem)

	for i := 0; i < 10; i++ {
		_, ok := <-sem
		fmt.Println("接收到信号量", ok)
	}

	fmt.Println("现在信号量的长度为:", len(sem))
}
