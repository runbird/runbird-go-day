package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type T111 struct {
	V int
}

func (t *T111) Incr(wg *sync.WaitGroup) {
	t.V++
	wg.Done()
}

func (t *T111) Print() {
	time.Sleep(1)
	fmt.Print(t.V)
}

func first_111() {
	var wg sync.WaitGroup
	wg.Add(10)
	var ts = make([]T111, 10)
	for i := 0; i < 10; i++ {
		ts[i] = T111{i}
	}
	for _, t := range ts {
		go t.Incr(&wg)
	}
	wg.Wait()
	// for range 循环里的变量 t 是临时变量。
	for _, t := range ts {
		go t.Print()
	}
	time.Sleep(5 * time.Second)
}

//--------------------------------
const N111 = 26

//下面的代码可以随机输出大小写字母，尝试在 A 处添加一行代码使得字母先按大写再按小写的顺序输出。
func second_111() {
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2 * N111)
	for i := 0; i < N111; i++ {
		go func(i int) {
			defer wg.Done()
			//A runtime.Gosched()
			runtime.Gosched()
			fmt.Printf("%c", 'a'+i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c", 'A'+i)
		}(i)
	}
	wg.Wait()
}
func main() {
	second_111()
}
