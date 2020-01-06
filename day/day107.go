package main

import (
	"fmt"
	"runtime"
	"time"
)

////panic: close of nil channel   ch未初始化,关闭时会报错
func first_107() {
	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}

func fix_first_107() {
	//var ch chan int
	ch := make(chan int, 1)
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		//  goroutine2 不延迟会导致提前关闭
		time.Sleep(time.Second)
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}

//--------------------

//程序执行到第二个 groutine 时，ch 还未初始化，导致第二个 goroutine 阻塞。
// 需要注意的是第一个 goroutine 不会阻塞。
func second_107() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

func fix_second_107() {
	//var ch chan int
	ch := make(chan int, 1)
	go func() {
		//ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

func main() {
	//first_107() //panic: close of nil channel   ch未初始化,关闭时会报错
	//second_107()
	//fix_first_107()
	fix_second_107()
}
