// clock_server
package main

import (
	"fmt"
	"net"
	"time"
)

func tserver() {
	protocol := "tcp"
	listen_Addr := "127.0.0.1:8888"
	listener, err := net.Listen(protocol, listen_Addr)
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Op == "accept" && opErr.Err.Error() == "use of closed network connection" {
				fmt.Println("Listener closed, stopping Accept loop")
				break
			}
			fmt.Println(err)
			continue
		}
		go func(c net.Conn) {
			fmt.Println("客户端地址：", conn.RemoteAddr())

			time.Sleep(10 * time.Second)
			fmt.Fprintln(conn, time.Now().Format("2006-01-02 15:04:05")) //将后面时间传给conon，相当于一个*write io

			defer conn.Close()
		}(conn)

	}
	err = listener.Close()
}
