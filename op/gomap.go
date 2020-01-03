package main

import "fmt"

func main() {
	// m := make(map[keyType]valueType)
	month := make(map[string]int)
	month["January"] = 1
	month["February"] = 2
	month["March"] = 3

	month2 := map[string]int{"January": 1, "February": 2, "March": 3}

	// 还可以写成这样
	month3 := map[string]int{
		"January":  1,
		"February": 2,
		"March":    3,
	}
	fmt.Println(month2, month3)

	month4 := map[string]int{}
	fmt.Println(month4) // 输出：map[]

	var month5 map[string]int
	fmt.Println(month5 == nil) // 输出：true

	match, exists := month["Match"]
	fmt.Println(match, exists) // 输出：0 false

	//delete函数 删除一个不存在的键值对时，delete函数不会报错，没任何作用。
	delete(month3, "January")

	//返回值的顺序有可能是不一样的，也就是说Map的遍历是无序的。
	for key, value := range month3 {
		fmt.Println(key, "=>", value)
	}

	fmt.Println("len(month3)=", len(month3))
}

func copy_map() {
	//Map是一种引用类型
	//Map是对底层数据的引用。编写代码的过程中，会涉及到Map拷贝、函数间传递Map等。
	// 跟Slice类似，Map指向的底层数据是不会发生copy的。
	m := map[string]int{
		"January":  1,
		"February": 2,
		"March":    3,
	}
	month := m
	delete(month, "February")
	fmt.Println(m)
	fmt.Println(month)
	//输出：
	//
	//map[January:1 March:3]
	//map[January:1 March:3]

	//上面的代码，将Map m赋值给month，删除了month中的一个键值对，m也发生了改变，说明Map拷贝时，m和month是共享底层数据的，
	// 改变其中一方数据，另一方也会随之改变。类似，在函数间传递Map时，其实传递的是Map的引用，不会涉及底层数据的拷贝，
	// 如果在被调用函数中修改了Map，在调用函数中也会感知到Map的变化。

	month2 := map[string]int{}
	m2 := map[string]int{
		"January":  1,
		"February": 2,
		"March":    3,
	}
	for key, value := range m2 {
		month2[key] = value
	}
	delete(month, "February")
	fmt.Println(m2)
	fmt.Println(month2)
	//输出：
	//
	//map[January:1 February:2 March:3]
	//map[January:1 March:3]
}
