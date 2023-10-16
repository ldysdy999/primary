package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Name string
}

func (f *Foo) SayHello() {
	fmt.Println("Hello from", f.Name)
}

func InvokeSayHello(obj interface{}) {
	// 获取值的反射表示
	value := reflect.ValueOf(obj)

	// 查找SayHello方法
	method := value.MethodByName("SayHello")

	// 检查方法是否存在
	if method.IsValid() {
		// 调用方法
		method.Call(nil)
	} else {
		fmt.Println("Method not found")
	}
}

func main() {
	foo := &Foo{Name: "John"}

	// 使用反射调用方法
	InvokeSayHello(foo)
}
