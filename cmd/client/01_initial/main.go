package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"
)

const maxBufferSize = 1024

func client(ctx context.Context, address string) (conn *net.UDPConn, err error) {
	raddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return
	}

	conn, err = net.DialUDP("udp", nil, raddr)
	if err != nil {
		return
	}
}

func main() {
	ctx := context.Background()
	conn, err := client(ctx, "127.0.0.1:3000")
	defer conn.Close()
	// reader io.Reader

	// create for loop
	//   - encode data
	//   - create reader
	//   - io.Copy to send request

	doneChan := make(chan error, 1)

	// go func() {
	// 	n, err := io.Copy(conn, reader)
	// 	conn.Write()
	// 	if err != nil {
	// 		doneChan <- err
	// 		return
	// 	}
	//
	// 	fmt.Printf("packet-written: bytes=%d\n", n)
	//
	// 	doneChan <- nil
	// }()
	//
	// select {
	// case <-ctx.Done():
	// 	fmt.Println("cancelled")
	// 	err = ctx.Err()
	// case err = <-doneChan:
	// }

	return
}
