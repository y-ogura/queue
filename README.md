# queue
The queue is a queue processing library for golang.

## Overview

## Getting Started

### Install
```bash
go get -u github.com/y-ogura/queue
```

### Quick Start

```go
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/y-ogura/queue"
)

var (
	maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
	maxQueueSize = flag.Int("max_queue_size", 10000, "The size of job queue")
)

func main() {
	flag.Parse()
	// create the job queue.
	jobQueue := make(chan queue.Job, *maxQueueSize)

	// start the dispatcher.
	dispatcher := queue.NewDispatcher(jobQueue, *maxWorkers)
	dispatcher.Run()

	for i := 0; i < 100; i++ {
		num := i
		// add new job
		jobQueue <- func() error {
			err := alertOdd(num)
			return err
		}
	}

	time.Sleep(10 * time.Second)
}

func alertOdd(num int) error {
	if num%2 == 0 {
		return nil
	}
	return fmt.Errorf("number %d is odd", num)
}
```
