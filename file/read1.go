//将n++移动到了if中，需要将if判断在添加一个条件

package main

import (
	"fmt"
	"io"
)

func (z *zimuguolv) Read1(p []byte) (int, error) {
	//当前位置>=字符串的长度，说明已经读取到结尾，返回EOF
	if z.cur >= len(z.src) {
		return 0, io.EOF
	}

	//定义一个剩余还没读到的长度
	x := len(z.src) - z.cur
	//n是本次遍历bound的索引，bound是本次读取的长度
	n, bound := 0, 0
	if x >= len(p) {
		//剩余长度超多缓冲区大小，说明本次可以完全填满缓冲区
		bound = len(p)
	} else {
		//剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区填不满
		bound = x
	}
	buf := make([]byte, bound)
	for n < bound && z.cur < len(z.src) {
		if char := alpha(byte(z.src[z.cur])); char != 0 {
			buf[n] = char
			// 将 n++ 移动到 if 语句内
			n++
		}
		// fmt.Printf("当前索引是%d,值是%v\n", n, string(buf[n]))
		//位置++
		z.cur++
	}
	fmt.Println(string(buf))
	copy(p, buf)
	return n, nil
}
