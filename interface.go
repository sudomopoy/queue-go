package queue

import "github.com/sudomopoy/go-queue/job"

type IQueue interface {
	Publish(j job.Job)
	Subscribe() job.Job
}
