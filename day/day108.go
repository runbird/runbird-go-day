package main

import (
	"fmt"
	"sync"
)

//sync.Map 没有 Len() 方法。
func fist_108() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	fmt.Println(m) //{{0 0} {{map[] true}} map[] 0}
}

//append() 并不是并发安全的
func second_108() {
	var wg sync.WaitGroup
	wg.Add(2)
	ints := make([]int, 0, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(len(ints))
}

func main() {
	fist_108()
	second_108() //可能小于2000
}
