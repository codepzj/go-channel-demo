package main

import (
	"fmt"
	"time"
)

func SelectCaseFunc(ch chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Println("接收到值为:", val)
			return
		case ch <- 666:
			fmt.Println("channel往里面置入了值")
		default:
			fmt.Println("上面channel都不位于就绪态, 暂时跳过")
		}
		time.Sleep(time.Millisecond * 500)
	}

}

func main() {
	ch := make(chan int, 1)
	SelectCaseFunc(ch)
}
