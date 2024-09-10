package queue

import "ropobackend/internal/queue/job"

type IQueue interface {
	Publish(j job.Job)
	Subscribe() job.Job
}
