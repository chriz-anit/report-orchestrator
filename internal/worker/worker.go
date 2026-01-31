package worker

import "report-orchestrator/internal/job"

type Worker interface {
	ProcessJob(job job.Job) error
}
