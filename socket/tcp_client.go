// clock_client
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func tclient() {
	protocol := "tcp"
	listen_Addr := "127.0.0.1:8888"
	conn, err := net.Dial(protocol, listen_Addr)
	if err != nil {
		fmt.Println(err)
	}
	start := time.Now()
	reader := bufio.NewReader(conn)
	fmt.Println(reader.ReadString('\n'))

	conn.Close()
	fmt.Println(time.Now().Sub(start))
}
