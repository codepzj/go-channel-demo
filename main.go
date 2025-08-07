package main

// import (
// 	"fmt"
// 	"time"
// )

// type EventChan struct {
// 	Name     string
// 	ChanList []chan int
// }

// type Chans []EventChan

// func (c Chans) Publish(name string, value int) {
// 	for _, ch := range c {
// 		// 找到订阅者
// 		if ch.Name == name {
// 			// 对所有订阅管道输送值
// 			for _, v := range ch.ChanList {
// 				v <- value
// 			}
// 		}
// 	}
// }

// func (c *Chans) Subscribe(name string) chan int {
// 	v := make(chan int, 1)
// 	find := false
// 	for i := range *c {
// 		// 如果找到
// 		if (*c)[i].Name == name {
// 			(*c)[i].ChanList = append((*c)[i].ChanList, v)
// 			find = true
// 			break
// 		}
// 	}
// 	// 没找到手动添加
// 	if !find {
// 		*c = append(*c, EventChan{
// 			Name:     name,
// 			ChanList: []chan int{v},
// 		})
// 	}
// 	return v

// }

// func Fn1(c *Chans) {
// 	go func() {
// 		ch1 := c.Subscribe("getNum")
// 		for val := range ch1 {
// 			fmt.Println("val1", val)
// 		}
// 	}()
// }
// func Fn2(c *Chans) {
// 	go func() {
// 		ch2 := c.Subscribe("getNum")
// 		for val := range ch2 {
// 			fmt.Println("val2", val)
// 		}
// 	}()
// }

// func main() {
// 	c := new(Chans)
// 	Fn1(c)
// 	Fn2(c)
// 	time.Sleep(60 * time.Millisecond)

// 	c.Publish("getNum", 1)
// 	time.Sleep(60 * time.Second)

// }
