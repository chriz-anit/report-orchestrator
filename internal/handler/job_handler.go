package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"report-orchestrator/internal/job"
	"report-orchestrator/internal/store"
)

type JobHandler struct {
	jobStore store.JobStore
}

func NewJobHandler(jobStore store.JobStore) *JobHandler {
	return &JobHandler{jobStore: jobStore}
}

type CreateJobRequest struct {
	ReportType string `json:"report_type"`
}

type CreateJobResponse struct {
	JobID  string        `json:"job_id"`
	Status job.JobStatus `json:"status"`
}

// CreateJob handles the creation of a new job based on the provided report type.
// It validates the request body, creates a new job with a unique ID and initial status, and stores it in the job store.
func (h *JobHandler) CreateJob(c *gin.Context) {
	var req CreateJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	req.ReportType = strings.TrimSpace(req.ReportType)
	if req.ReportType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "report_type is required",
		})
		return
	}

	now := time.Now()
	newJob := job.Job{
		ID:         uuid.NewString(),
		ReportType: req.ReportType,
		Status:     job.JobStatusRequested,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := h.jobStore.CreateJob(newJob); err != nil {
		switch err {
		case store.ErrJobAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{
				"error": "job with the same ID already exists",
			})
			return
		case store.ErrInvalidJobInput:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid job input",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create job",
			})
			return
		}
	}

	resp := CreateJobResponse{
		JobID:  newJob.ID,
		Status: newJob.Status,
	}

	c.JSON(http.StatusCreated, resp)
}

// GetJobByID retrieves a job by its ID from the job store and returns it in the response.
func (h *JobHandler) GetJobByID(c *gin.Context) {
	id := c.Param("id")
	id = strings.TrimSpace(id)

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	j, err := h.jobStore.GetJobByID(id)
	if err != nil {
		if err == store.ErrJobNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "job not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve job",
		})
		return
	}

	c.JSON(http.StatusOK, j)
}

// ListJobsByStatus retrieves a list of jobs filtered by their status from the job store and returns them in the response.
func (h *JobHandler) ListJobsByStatus(c *gin.Context) {
	status := c.Query("status")
	status = strings.TrimSpace(status)

	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "status query parameter is required",
		})
		return
	}

	parsedStatus, err := job.ParseJobStatus(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid job status",
		})
		return
	}

	jobs, err := h.jobStore.ListJobByStatus(job.JobStatus(parsedStatus))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list jobs",
		})
		return
	}

	c.JSON(http.StatusOK, jobs)
}
