package worker

import (
	"ropobackend/internal/queue"
	"time"
)

type Worker struct {
	Queue queue.IQueue
	Delay time.Duration
}

func NewWorker(queue queue.IQueue, delay time.Duration) Worker {
	return Worker{Queue: queue, Delay: delay}
}

func (w *Worker) Run() {
	for {
		job := w.Queue.Subscribe()
		job.Do()
		<-time.After(w.Delay)
	}
}
