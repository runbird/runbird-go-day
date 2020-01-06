package main

import (
	"fmt"
	"sync"
)

var c = make(chan int)
var a int

func f106() {
	a = 1
	<-c
}

//能正确输出1，不过主协程会阻塞 f106() 函数的执行。
func first_106() {
	go f106()
	c <- 0
	print(a)
}

//-------------------

type MyMutex_106 struct {
	count int
	sync.Mutex
}

//加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁。
func second_106() {
	var mu MyMutex_106
	mu.Lock()
	var mu1 = mu
	mu.count++
	mu.Unlock()
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)
}

func main() {
	//	first_106()
	second_106()
}
