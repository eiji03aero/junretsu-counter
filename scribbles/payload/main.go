package main

import (
	"fmt"

	"github.com/eiji03aero/junretsu-counter/internal/payload"
)

func main() {
	payloadBytes := payload.Encode(8)

	payloadStruct := payload.Payload{}
	payload.Decode(&payloadStruct, payloadBytes)

	fmt.Printf("decoded: %v\n", payloadStruct)
}
