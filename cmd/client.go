package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Client start!")

	err := testClient(8888)
	if err != nil {
		fmt.Println("Error!", err)
	}

	fmt.Println("done")
}

func testClient(p int) error {

	addr := fmt.Sprintf("localhost:%d", p)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}

	log.Println("Dialing")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	log.Println("Connected!")

	log.Println("Send hello")
	n, err := conn.Write([]byte("Hello!"))
	if err != nil {
		return err
	}
	fmt.Printf("XXX n = %d\n", n)

	time.Sleep(1 * time.Second)

	buf := make([]byte, 10)
	log.Println("Receive data")
	n, err = conn.Read(buf)
	if err != nil {
		return err
	}
	fmt.Printf("XXX buf = %s\n", string(buf[:n]))
	return nil
}
