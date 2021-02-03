package main

import (
	"log"
	"time"
)

type Work struct {
	ID  int
	Job string
}

type Worker struct {
	ID            int
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case work := <-w.Channel:
				log.Printf("id: %d, job: %s", work.ID, work.Job)
				time.Sleep(time.Duration(10*len(work.Job)) * time.Millisecond)
			case <-w.End:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.End <- true
}

type Collector struct {
	Work          chan Work
	End           chan bool
	WorkerChannel chan chan Work
	workers       []Worker
}

func NewCollector() *Collector {
	return &Collector{
		Work:          make(chan Work),
		End:           make(chan bool),
		WorkerChannel: make(chan chan Work),
		workers:       []Worker{},
	}
}

func (c *Collector) Start(workerCount int) {
	var i int

	for i < workerCount {
		i++
		log.Println("starting worker: ", i)
		worker := Worker{
			ID:            i,
			Channel:       make(chan Work),
			WorkerChannel: c.WorkerChannel,
			End:           make(chan bool),
		}
		worker.Start()
		c.workers = append(c.workers, worker)
	}

	go func() {
		for {
			select {
			case <-c.End:
				for _, w := range c.workers {
					w.Stop()
				}
			case work := <-c.Work:
				worker := <-c.WorkerChannel
				worker <- work
			}
		}
	}()
}

func main() {
	log.Println("starting")
	collector := NewCollector()
	collector.Start(2)

	works := []Work{
		{ID: 1, Job: "hoge desu ne"},
		{ID: 2, Job: "kore"},
		{ID: 3, Job: "sou datta noka to kiduita toki ha ososhi"},
		{ID: 4, Job: "kore ha naga ku naru bunsyou"},
		{ID: 5, Job: "hoge"},
	}

	for _, work := range works {
		log.Println("work: ", work.ID)
		collector.Work <- work
	}

	time.Sleep(1000 * time.Millisecond)
}
