package job

import (
	"math/rand"
	"time"
)

type Job struct {
	ID        int64
	fn        func() error
	Attempts  uint
	Sleep     time.Duration
	isStarted bool
	isDone    bool
	Timeout   time.Duration
	createdAt time.Time
	startedAt time.Time
	endAt     time.Time
}

func NewJob(fn func() error) *Job {
	return &Job{
		ID:        rand.Int63(),
		createdAt: time.Now(),
		isDone:    false,
		isStarted: false,
		fn:        fn,
	}
}
func (job *Job) WithSleep(sleep time.Duration) Job {
	job.Sleep = sleep
	return *job
}
func (job *Job) WithRetry(attempts uint) Job {
	job.Attempts = attempts
	return *job
}
func (job *Job) Do() {
	job.startedAt = time.Now()
	job.isStarted = true
	err := job.RunWithRetry()
	if err != nil {
		job.isDone = false
	}
	job.isDone = true
	job.endAt = time.Now()
}

func (job *Job) RunWithRetry() error {
	if err := job.fn(); err != nil {
		if job.Attempts--; job.Attempts > 0 {
			if job.Sleep == 0 {
				job.Sleep = time.Second
			}
			time.Sleep(job.Sleep)
			// double sleep time to avoid fast looping
			job.Sleep *= 2
			return job.RunWithRetry()
		}
		return err
	}
	return nil
}
