package store

import (
	"sync"

	"report-orchestrator/internal/job"
)

type InMemoryJobStore struct {
	mu sync.RWMutex
	jobs map[string]job.Job
}

func NewInMemoryJobStore() JobStore {
	return &InMemoryJobStore{
		jobs: make(map[string]job.Job),
	}
}

func (s *InMemoryJobStore) CreateJob(j job.Job) error {
	return nil
}

func (s *InMemoryJobStore) GetJobByID(id string) (job.Job, error) {
	return job.Job{}, nil
}

func (s *InMemoryJobStore) UpdateJob(j job.Job) error {
	return nil
}

func (s *InMemoryJobStore) ListJobByStatus(status job.JobStatus) ([]job.Job, error) {
	return []job.Job{}, nil
}