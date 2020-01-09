package main

import (
	"fmt"
	"sort"
	"time"
)

//panic 写成开启还未来得及执行，chan就close
func first_110() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("Ok!")
	time.Sleep(time.Second * 20)
}

//-------------
type S110 struct {
	v int
}

func second_110() {
	s := []S110{{1}, {3}, {2}, {5}}
	//A 一行代码实现S按升序排序
	sort.Slice(s, func(i, j int) bool {
		return s[i].v < s[j].v
	})
	fmt.Printf("%#v", s)
}

func main() {
	second_110()
}
