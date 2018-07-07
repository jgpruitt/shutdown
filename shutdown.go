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

package shutdown

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// A Task is a function that needs to be run before the process exits.
type Task func()

var (
	once  sync.Once
	mu    sync.Mutex
	tasks []Task
	exit  = os.Exit
	block = func() {
		select{}
	}
)

// Add registers a Task to be executed just before the process exits.
func Add(t Task) {
	mu.Lock()
	defer mu.Unlock()
	tasks = append(tasks, t)
}

// OnSignal blocks until SIGINT, SIGHUP, or SIGTERM is received.
// Then, it runs all shutdown Tasks in FILO order.
// Finally, it exits the process with exit code = 0.
func OnSignal() {
	var c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	<-c
	Now(0)
}

// Now initiates the shutdown process immediately.
// Then, it runs all shutdown Tasks in FILO order.
// Finally, it exits the process with exit code = "code".
func Now(code int) {
	once.Do(func() {
		mu.Lock()
		defer mu.Unlock() // not that it should matter
		for i := len(tasks) - 1; i >= 0; i-- {
			tasks[i]()
		}
		exit(code)
	})
	// if multiple goroutines try to shutdown concurrently, block the ones that
	// did not make it into the once.Do call since they are expected a call to
	// Now to exit the app (i.e. never return)
	block()
}
