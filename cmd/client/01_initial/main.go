package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/eiji03aero/junretsu-counter/internal/payload"
)

func client(ctx context.Context, address string) (conn *net.UDPConn, err error) {
	raddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return
	}

	conn, err = net.DialUDP("udp", nil, raddr)
	return
}

func main() {
	ctx := context.Background()

	conn, err := client(ctx, "127.0.0.1:3000")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	doneChan := make(chan error, 1)

	go func() {
		send := func (buttonId uint8) error {
			data := payload.Encode(buttonId)
			buf := bytes.NewBuffer(data)
			_, err := io.Copy(conn, buf)
			return err
		}

		timeout := time.After(1 * time.Second)
LOOP:
		for {
			select {
			case <-timeout:
				break LOOP
			default:
				send(0)
			}
		}

		send(1)

		doneChan <- nil
	}()

	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
		err = ctx.Err()
		panic(err)
	case <-doneChan:
		fmt.Println("done")
	}
}
