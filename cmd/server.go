package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Server start!")

	err := testServer(8888)
	if err != nil {
		fmt.Println("Error!", err)
	}

	fmt.Println("done")
}

func testServer(p int) error {
	addr := fmt.Sprintf("0.0.0.0:%d", p)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	log.Println("Waiting..")
	conn, err := l.Accept()
	if err != nil {
		return err
	}
	log.Println("Accepted!")

	log.Println("Send Konichiwa")
	n, err := conn.Write([]byte("Konichiwa"))
	if err != nil {
		return err
	}
	fmt.Printf("XXX n = %d\n", n)

	time.Sleep(2 * time.Second)

	buf := make([]byte, 10)
	log.Println("Receive data")
	n, err = conn.Read(buf)
	if err != nil {
		return nil
	}
	fmt.Printf("XXX buf = %s\n", string(buf[:n]))

	return nil
}
