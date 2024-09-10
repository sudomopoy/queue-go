package job

import (
	"fmt"
	"github.com/google/uuid"
	"ropobackend/internal/service"
	"ropobackend/logger"
	"time"
)

type Job struct {
	ID        string
	Service   service.Service
	fn        func() error
	Attempts  uint
	Sleep     time.Duration
	isStarted bool
	isDone    bool
	Timeout   time.Duration
	createdAt time.Time
	startedAt time.Time
	endAt     time.Time
	Logger    logger.Logger
}

func NewJob(fn func() error, lg logger.Logger) *Job {
	return &Job{
		ID:        uuid.New().String(),
		createdAt: time.Now(),
		isDone:    false,
		isStarted: false,
		fn:        fn,
		Logger:    lg,
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
		job.Logger.WithField(fmt.Sprintf("job/%v", job.ID), err).Error(fmt.Sprintf("job %v fail", job.ID))
	}
	job.isDone = true
	job.endAt = time.Now()
}

func (job *Job) RunWithRetry() error {
	if err := job.fn(); err != nil {
		if job.Attempts--; job.Attempts > 0 {
			job.Logger.WithField(fmt.Sprintf("job/%v", job.ID), err).Error(fmt.Sprintf("job %d attempt fail", job.Attempts))
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
