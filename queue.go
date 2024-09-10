package queue

import (
	"time"

	"github.com/sudomopoy/go-queue/job"
	"github.com/sudomopoy/go-queue/worker"
)

type Queue struct {
	Jobs    chan job.Job
	Workers []worker.Worker
}

func NewQueue(queueCap int) *Queue {
	jobs := make(chan job.Job, queueCap)
	workers := make([]worker.Worker, 0)
	return &Queue{Jobs: jobs, Workers: workers}
}

var _ IQueue = &Queue{}

func (q *Queue) AddWorkers(num int, workersDelay time.Duration) {
	var workers []worker.Worker
	for i := 0; i < num; i++ {
		workers = append(workers, worker.NewWorker(q, workersDelay))
	}
	q.Workers = workers
}

func (q *Queue) RunWorkers() {
	for _, w := range q.Workers {
		go w.Run()
	}
}

func (q *Queue) Publish(j job.Job) {
	go func() {
		<-time.After(j.Sleep)
		q.Jobs <- j
	}()
}

func (q *Queue) Subscribe() job.Job {
	return <-q.Jobs
}
