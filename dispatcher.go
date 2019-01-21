package queue

import "log"

type dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
}

// NewDispatcher creates, and returns a new dispatcher
func NewDispatcher(jobQueue chan Job, maxWorkers int) Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)

	return &dispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
	}
}

// Dispatcher dispatcher interface
type Dispatcher interface {
	Run()
}

// Run dispatcher run with main
func (d *dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i+1, d.workerPool)
		worker.start()
	}

	go d.dispatch()
}

func (d *dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func() {
				log.Printf("[queue] fetching workerJobQueue\n")
				workerJobQueue := <-d.workerPool
				log.Printf("[queue] adding job to workerJobQueue\n")
				workerJobQueue <- job
			}()
		}
	}
}
