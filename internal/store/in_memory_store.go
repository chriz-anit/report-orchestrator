package store

import (
	"sync"
	"time"

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
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.jobs[j.ID]; exists {
		return ErrJobAlreadyExists
	}

	if j.Status != job.JobStatusRequested {
		return ErrInvalidJobStatus
	}

	now := time.Now()
	j.CreatedAt = now
	j.UpdatedAt = now

	s.jobs[j.ID] = j
	return nil
}

func (s *InMemoryJobStore) GetJobByID(id string) (job.Job, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	j, exists := s.jobs[id]
	if !exists {
		return job.Job{}, ErrJobNotFound
	}

	return j, nil
}

func (s *InMemoryJobStore) UpdateJob(j job.Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	existingJob, exists := s.jobs[j.ID]
	if !exists {
		return ErrJobNotFound
	}

	if !job.IsValidTransition(existingJob.Status, j.Status) {
		return ErrInvalidJobStatus
	}

	j.CreatedAt = existingJob.CreatedAt
	j.UpdatedAt = time.Now()

	s.jobs[j.ID] = j
	return nil
}

func (s *InMemoryJobStore) ListJobByStatus(status job.JobStatus) ([]job.Job, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []job.Job
	for _, j := range s.jobs {
		if j.Status == status {
			result = append(result, j)
		}
	}
	return result, nil
}