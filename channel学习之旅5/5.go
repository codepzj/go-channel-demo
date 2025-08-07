package main

import (
	"fmt"
	"sync"
	"time"
)

type EventChans struct {
	query []chan int
	mu    sync.RWMutex
}

type Eventbus struct {
	ec map[string]*EventChans // 设置为指针的目的是能够在订阅和取消订阅的时候原地修改
	mu sync.RWMutex
}

func NewEventbus() *Eventbus {
	return &Eventbus{
		mu: sync.RWMutex{},
		ec: make(map[string]*EventChans, 10),
	}
}

type IEventbus interface {
	Subscribe(event string, ch chan int)
	UnSubscribe(event string, ch chan int)
	Publish(event string, message int)
}

func (eb *Eventbus) Subscribe(event string, ch chan int) chan int {
	// 对map加互斥锁
	eb.mu.Lock()
	ec, ok := eb.ec[event]
	if !ok {
		ec = &EventChans{}
		eb.ec[event] = ec // 初始化
	}
	eb.mu.Unlock()

	ec.mu.Lock()
	defer ec.mu.Unlock()
	ec.query = append(ec.query, ch)
	return ch // 返回用于监听广播的值
}

func (eb *Eventbus) UnSubscribe(event string, ch chan int) {
	// 对map加互斥锁
	eb.mu.Lock()
	ec, ok := eb.ec[event]
	eb.mu.Unlock()
	if !ok {
		return // 没找到直接返回
	}

	ec.mu.Lock()
	defer ec.mu.Unlock()
	for i, ech := range ec.query {
		if ch == ech {
			ec.query = append(ec.query[0:i], ec.query[i+1:]...)
			break
		}
	}
}

func (eb *Eventbus) Publish(event string, message int) {
	// 对map加读锁
	eb.mu.RLock()
	ec, ok := eb.ec[event]
	eb.mu.RUnlock()
	if !ok {
		return // 没找到直接返回
	}

	// 对切片加读锁
	ec.mu.RLock()
	defer ec.mu.RUnlock()
	for _, ch := range ec.query {
		ch <- message
	}
}

func subscribe(eb *Eventbus) {
	ch := make(chan int, 1)
	eb.Subscribe("event1", ch)
	wg.Done()
	fmt.Println("订阅完毕，开始等待接收")
	fmt.Println("接收到值:", <-ch)
}

var wg sync.WaitGroup

func main() {
	eb := NewEventbus()
	go subscribe(eb)

	// 发布事件

	wg.Add(1)
	wg.Wait() // 等待订阅成功
	eb.Publish("event1", 666)
	time.Sleep(time.Second)
}
