package store

import (
	"report-orchestrator/internal/job"
)

type JobStore interface {
	CreateJob(job job.Job) error
	GetJobByID(id string) (job.Job, error)
	UpdateJob(job job.Job) error
	ListJobByStatus(status job.JobStatus) ([]job.Job, error)
}
