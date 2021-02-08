package main

import (
	"fmt"
	"net"

	"github.com/eiji03aero/junretsu-counter/internal/payload"
)

const maxBufferSize = 64

func main() {
	address := "127.0.0.1:3000"

	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		return
	}

	defer pc.Close()

	doneChan := make(chan error, 1)
	buffer := make([]byte, maxBufferSize)
	count := 0

	go func() {
		for {
			_, _, err := pc.ReadFrom(buffer)
			if err != nil {
				doneChan <- err
				return
			}

			payl := payload.Payload{}
			err = payload.Decode(&payl, buffer)
			if err != nil {
				doneChan <- err
				return
			}

			if payl.ButtonId == 1 {
				doneChan <- nil
				return
			} else {
				count++
			}
		}
	}()

	select {
	case err = <-doneChan:
		fmt.Println("done: ", count)
	}

	return
}
