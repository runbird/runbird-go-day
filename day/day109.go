package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

//按照 go 的语法，小写开头的方法、属性或 struct 是私有的，
// 同样，在 json 解码或转码的时候也无法实现私有属性的转换。
//私有属性 name 也不应该加 json 的标签。
type Pepople struct {
	name string `json:"name"`
}

func first_109() {
	js := `{"name":"11"}`
	var p Pepople
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}

//-----------------------
var ip string
var port int

//补充 A、B 两处代码，实现程序能解析 ip 和 prot 参数，默认值是  0.0.0.0 和 8000。
func init() {
	//A
	flag.StringVar(&ip, "ip", "0.0.0.0", "ip address")
	flag.IntVar(&port, "port", 8000, "port number")
	//B
}

func main() {
	first_109()
	flag.Parse()
	fmt.Printf("%s,%d", ip, port)
}
