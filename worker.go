package queue

import "log"

// Job holds the attributes needed to perform unit of work
type Job func() error

// Worker worker struct
type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

// NewWorker create takes a numeric id and a channel w/ worker pool
func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

func (w Worker) start() {
	go func() {
		for {
			// Add my jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// Dispatcher has added a job to my jobQueue.
				log.Printf("[queue] worker%d: started\n", w.id)
				err := job()
				if err != nil {
					log.Printf("[queue] Error worker_id: %d, message: %v\n", w.id, err)
				}
				log.Printf("[queue] worker%d: completed!\n", w.id)
			case <-w.quitChan:
				// We have been asked to stop.
				log.Printf("[queue] worker%d: stopping\n", w.id)
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}
