package main

import (
	"fmt"
	"time"
)

// select 的用法有点类似 switch 语句，但 select 不会有输入值 而且只用于信道操作。
// select 用于从多个发送或接收信道操作中进行选择，语句会阻塞直到其中有信道可以操作，如果有多个信道可以操作，会随机选择其中一个 case 执行。
func service1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from service1"
}

func service2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from service2"
}

func service3(ch chan string) {
	ch <- "from service3"
}
func service4(ch chan string) {
	ch <- "from service3"
}

//上面的例子执行到 select 语句的时候回发生阻塞，main 协程等待一个 case 操作可执行，
// 很明显是 service2 先准备好读取的数据（休眠 1s），所以输出 from service2。
func selectMethod() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go service1(ch1)
	go service2(ch2)

	// 若ch1 ch2无休眠时间，main线程 time.Sleep(2*time.Second) 则随机输出
	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}
}

//执行到 select 语句的时候，由于信道 ch1、ch2 都没准备好，直接执行 default 语句。
func selectDefault() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go service3(ch1)
	go service4(ch2)

	//  在 select 语句之前加了 1s 延时，等待 ch1 ch2 准备就绪。因为两个通道都准备好了，
	//  所以不会走 default 语句。随机输出 from service1 或 from service2。
	//	time.Sleep(time.Second)

	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	default:
		fmt.Println("no case ok")
	}
}

func selectNotice() {
	//信道的默认值是 nil，不能对 nil 信道进行读写操作
	//case 分支中如果信道是 nil，该分支就会被忽略，那么上面就变成空 select{} 语句,
	// 阻塞主协程，调度 service1 协程，在 nil 信道上操作，便报[chan send (nil chan)] 错误
	var ch chan string

	go service3(ch)
	select {
	case s1 := <-ch:
		fmt.Println(s1)
		//使用default避免nil操作
	default:
		fmt.Println("default")
	}
}

func selectTime() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go service1(ch1)
	go service2(ch1)

	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	//case s3:= <-time.After(1*time.Second):
	//	fmt.Println(s3)
	case <-time.After(1 * time.Second):
		fmt.Println("time out")
	}
}

func main() {
	selectMethod() //from service2
	selectDefault()
	selectNotice()
	selectTime()
}
