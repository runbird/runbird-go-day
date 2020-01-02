package main

import "fmt"

//编译通过。true 是预定义标识符可以用作变量名，但是不建议这么做
func first_100() {
	true := false
	fmt.Println(true)
}

//--------------------------
func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret := 10
		defer func() {
			ret += 1
		}()
	}
	return
}

// 知识点：变量作用域和defer 返回值。
// 可以查看文章《5 年 Gopher 都不知道的 defer 细节，你别再掉进坑里！》

func second_100() {
	result := watShadowDefer(50)
	fmt.Println(result) //100
}

func main() {
	first_100()
	second_100()
}
