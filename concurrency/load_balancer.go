// based rob
package main

import (
	"container/heap"
	"log"
	"time"
)

const (
	STATUS_ERROR     int = 1
	STATUS_QUEUED    int = 2
	STATUS_RUNNING   int = 4
	STATUS_CANCELLED int = 8
	STATUS_COMPLETED int = 16
)

type Request struct {
	Status    int
	Progress  chan int
	LastP     int
	Stop      chan struct{}
	Time      time.Time
	format    string
	infile    string
	outfile   string
	destdir   string
	errString string
}

type Worker struct {
	requests chan *Request
	pending  int
	index    int
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests
		req.Time = time.Now()
		if req.Status == STATUS_QUEUED {
			req.Status = STATUS_RUNNING
			req.Status = req.workFn()
			req.Time = time.Now()
		}
		done <- w
	}
}

type Pool []*Worker

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pool) Push(x interface{}) {
	*p = append(*p, x.(*Worker))
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan *Request) {
	for {
		select {
		case req, ok := <-work:
			if ok {
				b.dispatch(req)
			} else {
				log.Println("requests channel is closed, exiting...")
				return
			}
		case w := <-b.done:
			b.completed(w)
		case <-time.After(keep_time):
			cleanup()
		}
	}
}

func (b *Balancer) dispatch(req *Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- req
	w.pending++
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Worker) {
	w.pending--
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}

// the above was taken wholesale from
// Rob Pike's "Concurrency is not Parallelism"
//
// end load balancer stuff
