# queue
The queue is a queue processing library for golang.

[![Sourcegraph](https://sourcegraph.com/github.com/y-ogura/queue/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/y-ogura/queue?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/y-ogura/queue)
[![Go Report Card](https://goreportcard.com/badge/github.com/y-ogura/queue?style=flat-square)](https://goreportcard.com/report/github.com/y-ogura/queue)
[![Build Status](http://img.shields.io/travis/y-ogura/queue.svg?style=flat-square)](https://travis-ci.org/y-ogura/queue)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/y-ogura/queue/master/LICENSE)


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

## License
[MIT](https://github.com/y-ogura/queue/blob/master/LICENSE)
