package worker

import (
	"time"

	queue "github.com/sudomopoy/queue-go"
)

type Worker struct {
	Queue queue.Queue
	Delay time.Duration
}

func NewWorker(queue queue.Queue, delay time.Duration) Worker {
	return Worker{Queue: queue, Delay: delay}
}

func (w *Worker) Run() {
	for {
		job := w.Queue.Subscribe()
		job.Do()
		<-time.After(w.Delay)
	}
}
