package job

import (
	"errors"
	"strings"
)

var ErrInvalidJobStatus = errors.New("invalid job status")

func IsValidJobStatus(status string) bool {
	switch JobStatus(status) {
	case JobStatusRequested, JobStatusRunning, JobStatusCompleted, JobStatusFailed:
		return true
	default:
		return false
	}
}

func ParseJobStatus(status string) (JobStatus, error) {
	s := strings.TrimSpace(strings.ToUpper(status))
	
	if !IsValidJobStatus(s) {
		return "", ErrInvalidJobStatus
	}
	
	return JobStatus(s), nil
}