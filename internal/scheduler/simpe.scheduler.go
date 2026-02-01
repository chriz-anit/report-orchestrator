package scheduler

import (
	"report-orchestrator/internal/job"
	"report-orchestrator/internal/store"
	"report-orchestrator/internal/worker"
	"time"
)

// SimpleScheduler polls requested jobs at a fixed interval
// and dispatches them to a worker for processing.
type SimpleScheduler struct {
	jobStore store.JobStore
	worker   worker.Worker
	interval time.Duration

	ticker *time.Ticker
	stopCh chan struct{}
}

func NewSimpleScheduler(jobStore store.JobStore, worker worker.Worker, interval int) Scheduler {
	return &SimpleScheduler{
		jobStore: jobStore,
		worker:   worker,
		interval: time.Duration(interval) * time.Second,
		stopCh:   make(chan struct{}),
	}
}

func (ss *SimpleScheduler) Start() {
	ss.ticker = time.NewTicker(ss.interval)
	go func() {
		for {
			select {
			case <-ss.ticker.C:
				ss.dispatchJobs()
			case <-ss.stopCh:
				return
			}
		}
	}()
}

func (ss *SimpleScheduler) Stop() {
	if ss.ticker != nil {
		ss.ticker.Stop()
	}
	close(ss.stopCh)
}

func (ss *SimpleScheduler) dispatchJobs() {
	jobs, err := ss.jobStore.ListJobByStatus(job.JobStatusRequested)
	if err != nil {
		return
	}

	for _, j := range jobs {
		ss.worker.ProcessJob(j)
	}
}
