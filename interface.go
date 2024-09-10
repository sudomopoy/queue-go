package queue

import "github.com/sudomopoy/queue-go/job"

type Queue interface {
	Publish(j job.Job)
	Subscribe() job.Job
}
