# shutdown
A package for doing graceful shutdowns.

[![GoDoc](https://godoc.org/github.com/jgpruitt/shutdown?status.svg)](https://godoc.org/github.com/jgpruitt/shutdown)
[![Go Report Card](https://goreportcard.com/badge/github.com/jgpruitt/shutdown)](https://goreportcard.com/report/github.com/jgpruitt/shutdown)

## Installation

```sh
go get -u github.com/jgpruitt/shutdown
```

## Example

Register functions that need to be run on a graceful shutdown of the process.
They will be executed in last-in-first-out order.

```go
func ExampleAdd() {
	// register a task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe close a logging subsystem
		fmt.Println("I will be executed last.")
	})

	// register another task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe close database connections
		fmt.Println("I get executed second.")
	})

	// register a third task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe gracefully close an http.Server
		fmt.Println("I get executed first.")
	})
}
```

Use ```shutdown.OnSignal()``` to block until the process receives a signal to die.
This initiates a graceful shutdown. Shutdown tasks are executed and then the process exits.

```go
func ExampleOnSignal() {
	// register a task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe close a logging subsystem
		fmt.Println("I will be executed last.")
	})

	// register another task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe close database connections
		fmt.Println("I get executed second.")
	})

	// register a third task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe gracefully close an http.Server
		fmt.Println("I get executed first.")
	})

	// this blocks until a SIGINT, SIGHUP, or SIGTERM is received
	// then it runs the shutdown tasks in FILO order
	// then it exits the process
	shutdown.OnSignal()
}
```

Use ```shutdown.Now()``` to initiate an immediate graceful shutdown.

```go
func ExampleNow() {
	// register a task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe close a logging subsystem
		fmt.Println("I will be executed last.")
	})

	// register another task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe close database connections
		fmt.Println("I get executed second.")
	})

	// register a third task to be executed on shutdown
	shutdown.Add(func() {
		// in reality, maybe gracefully close an http.Server
		fmt.Println("I get executed first.")
	})

	// this runs the shutdown tasks in FILO order
	// then it exits the process with exit code = 99
	shutdown.Now(99)
}
```

All functions are safe to use concurrently.
