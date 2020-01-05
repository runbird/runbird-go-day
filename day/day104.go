//https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648466918&idx=2&sn=151a8135f22563b7b97bf01ff480497b&chksm=f2474389c530ca9f3dc2ae1124e4e5ed3db4c45096924265bccfcb8908a829b9207b0dd26047&scene=21#wechat_redirect
package main

import (
	"fmt"
	"sync"
	"time"
)

func first_104() {
	fmt.Println(doubleScore(0))
	fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))
}

//函数的 return value 不是原子操作，而是在编译器中分解为两部分：返回值赋值 和 return
func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	return source * 2
}

//-----------------------
var mu sync.RWMutex
var count int

//当写锁阻塞时，新的读锁是无法申请的（有效防止写锁饥饿），导致死锁。
func second_104() {
	go A_104()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}
func A_104() {
	mu.RLock()
	defer mu.RUnlock()
	B_104()
}
func B_104() {
	time.Sleep(5 * time.Second)
	C_104()
}
func C_104() {
	mu.RLock()
	defer mu.RUnlock()
}

func main() {
	//first_104()
	second_104() //fatal error: all goroutines are asleep - deadlock!
}
