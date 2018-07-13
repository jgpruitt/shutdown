// MIT License
//
// Copyright (c) 2018 John Pruitt
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package shutdown_test

import (
	"fmt"
	"github.com/jgpruitt/shutdown"
)

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
