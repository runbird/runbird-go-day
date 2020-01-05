package main

import (
	. "sync"
	"time"
)

func first_105() {
	//A. 当一个 goroutine 获得了 Mutex 后，其他 goroutine 就只能乖乖的等待，除非该 goroutine 释放这个 Mutex；
	//
	//B. RWMutex 在读锁占用的情况下，会阻止写，但不阻止读；
	//
	//C. RWMutex 在写锁占用情况下，会阻止任何其他 goroutine（无论读和写）进来，整个锁相当于由该 goroutine 独占；
	//
	//D. Lock() 操作需要保证有 Unlock() 或 RUnlock() 调用与之对应；
	//ABC
}

//WaitGroup 在调用 Wait() 之后不能再调用 Add() 方法的。
func second_105() {
	var wg WaitGroup //sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}
func main() {
	second_105() //panic: sync: WaitGroup is reused before previous Wait has returned

}
