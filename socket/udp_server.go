package main

import (
	"fmt"
	"net"
)

func userver() {
	addr := ":8888"
	protocol := "udp"
	packetconn, err := net.ListenPacket(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		ctx := make([]byte, 1024)
		n, client_addr, err := packetconn.ReadFrom(ctx)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("客户端：[%s]发送的消息：%s\n", client_addr, string(ctx[:n]))

		packetconn.WriteTo([]byte("哈哈哈哈"), client_addr)
	}
	packetconn.Close()
}
