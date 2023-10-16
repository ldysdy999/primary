//将JSON字符串转换回Go结构体。在这种情况下，解码过程是将JSON格式的文本表示转换回Go结构体的内存表示。

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := `{"name":"Alice","age":30}`
	var p Person
	json.Unmarshal([]byte(jsonData), &p)
	fmt.Println(p)
	fmt.Printf("%#v", p)
}
