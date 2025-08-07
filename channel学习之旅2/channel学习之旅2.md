## 缓冲区为 1 和无缓冲区的 channel 有什么区别

### 缓冲区为 1

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1

	fmt.Println("ch中的值为", <-ch)
}
```

### 无缓冲区

```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 1

	fmt.Println("ch中的值为", <-ch)
}
```

缓冲区为 1 能正常输出 1
而无缓冲区会死锁，因为在往 channel 置入值，当前 goroutine 就会被挂起，消费者无法接收 channel 传入的值
