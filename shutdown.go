package shutdown

import (
	"os"
	"sync"
	"os/signal"
)

type Task func()

type Shutdown struct {
	once sync.Once
	mu sync.Mutex
	tasks []Task
}

func (sd *Shutdown) Add(t Task) {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	sd.tasks = append(sd.tasks, t)
}

func (sd *Shutdown) OnSignal() {
	var c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	sd.Now(0)
}

func (sd *Shutdown) Now(code int) {
	sd.once.Do(func() {
		sd.mu.Lock()
		defer sd.mu.Unlock()
		for i := len(sd.tasks) - 1; i >= 0; i-- {
			sd.tasks[i]()
		}
		os.Exit(code)
	})
}
