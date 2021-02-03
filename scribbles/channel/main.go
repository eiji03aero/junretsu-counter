package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Payload struct {
	Worker int
	Work   int
}

var now = time.Now()

func (p *Payload) String() string {
	t := time.Now().Sub(now).Seconds()
	return fmt.Sprintf("time: %f, wkr: %d, w: %d", t, p.Worker, p.Work)
}

func makeRequests(pch chan Payload, wg *sync.WaitGroup, numWorkers int, numWorks int) {
	for i := 0; i < numWorkers; i++ {
		ii := i
		go func() {
			for j := 0; j < numWorks; j++ {
				p := Payload{Worker: ii, Work: j}
				log.Println("Sending: ", p.String())
				pch <- p
			}
			wg.Done()
		}()
	}
}

func handleRequests(pch chan Payload, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func() {
			for p := range pch {
				log.Println("Received: ", p.String())
				d := time.Duration((p.Worker + p.Work + 1) * 400)
				time.Sleep(time.Millisecond * d)
			}
		}()
	}
}

func test1() {
	var wg sync.WaitGroup
	pch := make(chan Payload)

	log.Println("starting unbuffered")
	wg.Add(2)
	go handleRequests(pch, 1)
	go makeRequests(pch, &wg, 2, 5)
	wg.Wait()
	log.Println("finishing unbuffered")
	log.Println("")
}

func test2() {
	var wg sync.WaitGroup
	pch := make(chan Payload, 20)

	log.Println("starting buffered")
	wg.Add(2)
	go handleRequests(pch, 1)
	go makeRequests(pch, &wg, 2, 5)
	wg.Wait()
	log.Println("finishing buffered")
	log.Println("")
}

func main() {
	test1()
	test2()
}
