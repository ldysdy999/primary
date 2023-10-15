package main

import (
	"io"
	"log"
)

type zimuguolv struct {
	src string //输入的字符串
	cur int    //读取到当前的位置
}

// 字符过滤函数
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func main() {
	zmreader := zimuguolv{
		src: "ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!ldy shi ni die 2022 go !!!!",
	}
	p := make([]byte, 30)
	for {
		n, err := zmreader.Read(p)
		if err == io.EOF {
			log.Printf("[eof错误]")
			break
		}
		log.Printf("[读取到的长度为%d 内容%s]", n, string(p[:n]))
	}
}
