package worker

import (
	"time"
	
	"report-orchestrator/internal/job"
	"report-orchestrator/internal/store"
)

type SimpleWorker struct {
	jobStore store.JobStore
}

func NewSimpleWorker(jobStore store.JobStore) Worker {
	return &SimpleWorker{
		jobStore: jobStore,
	}
}

func (sw *SimpleWorker) ProcessJob(j job.Job) error {
	if !job.IsValidTransition(j.Status, job.JobStatusRunning) {
		return store.ErrInvalidJobStatusTransition
	}

	j.Status = job.JobStatusRunning
	now := time.Now()
	j.UpdatedAt = now
	j.StartTime = now

	if err := sw.jobStore.UpdateJob(j); err != nil {
		return err
	}

	err := sw.executeJob(j)

	now = time.Now()
	j.UpdatedAt = now
	j.EndTime = now

	if err != nil {
		j.Status = job.JobStatusFailed
		j.ErrorMessage = err.Error()
	} else {
		j.Status = job.JobStatusCompleted
	}

	return sw.jobStore.UpdateJob(j)
}

func (sw *SimpleWorker) executeJob(j job.Job) error {
	// Replace this with actual job processing logic.
	time.Sleep(2 * time.Second)
	return nil
}
