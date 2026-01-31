package store

import (
	"errors"

	"report-orchestrator/internal/job"
)

type JobStore interface {
	CreateJob(job job.Job) error
	GetJobByID(id string) (job.Job, error)
	UpdateJob(job job.Job) error
	ListJobByStatus(status job.JobStatus) ([]job.Job, error)
}

var (
	ErrJobAlreadyExists = errors.New("job already exists")
	ErrJobNotFound      = errors.New("job not found")
	ErrInvalidJobStatus = errors.New("invalid job status")
)
