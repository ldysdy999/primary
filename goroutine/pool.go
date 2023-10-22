package main

import (
	"fmt"
	"sync"
)

func main() {
	intPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("调用了我，创建了一个该对象")
			return 1 //这个就是返回的对象
		},
	}

	a := intPool.Get()
	intPool.Put(a)     //放入一个int了类型对象
	b := intPool.Get() //取值为上一个放入的对象，不再调用New了
	c := intPool.Get()
	fmt.Println(a, b, c)
}
