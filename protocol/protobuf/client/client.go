package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	strIP := "127.0.0.1:6600"
	var conn net.Conn
	var err error

	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("connect", strIP, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect...")
	}
	fmt.Println("connect", strIP, "success")
	defer conn.Close()

	cnt := 0
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		cnt++
		stSend := &proto
	}

}
