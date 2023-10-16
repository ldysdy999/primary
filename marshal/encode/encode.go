//将Go结构体转换为JSON字符串。在这种情况下，编码过程是将Go结构体的内存表示转换为JSON格式的文本表示。

package marshal

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	jsonData, _ := json.MarshalIndent(p, "", "\t")
	fmt.Println(string(jsonData))
}
