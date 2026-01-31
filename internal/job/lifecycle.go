package job

// IsValidTransition checks if a transition from the current job status to the next
// is according to the defined lifecycle.
func IsValidTransition(current, next JobStatus) bool {
	switch current {
		case JobStatusRequested:
			return next == JobStatusRunning
		case JobStatusRunning:
			return next == JobStatusCompleted || next == JobStatusFailed
		default:
			return false
	}
}