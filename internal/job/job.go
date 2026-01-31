package job

import "time"

// Job represents a report generation job with its associated metadata and status.
type Job struct {
	ID         string
	ReportType string
	Status     JobStatus

	CreatedAt time.Time
	UpdatedAt time.Time

	StartTime time.Time
	EndTime   time.Time

	ErrorMessage string
}

// JobStatus represents the lifecycle status of a job.
type JobStatus string

const (
	JobStatusRequested JobStatus = "REQUESTED"
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusCompleted JobStatus = "COMPLETED"
	JobStatusFailed    JobStatus = "FAILED"
)
