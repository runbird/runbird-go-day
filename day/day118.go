package main

import (
	"sync"
)

//1.Go 语言中中大多数数据类型都可以转化为有效的 JSON 文本，下面几种类型除外。
//
//A. 指针
//B. channel
//C. complex
//D. 函数
// BCD
const N118 = 10

func second_118() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N118)
	for i := 0; i < N118; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}

func fix_second_118() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N118)
	for i := 0; i < N118; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
}

func main() {

}
