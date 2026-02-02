package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"report-orchestrator/internal/scheduler"
	"report-orchestrator/internal/store"
	"report-orchestrator/internal/worker"
)

const (
	serverPort               = ":8080"
	schedulerIntervalSeconds = 5
)

func main() {
	jobStore := store.NewInMemoryJobStore()
	jobWorker := worker.NewSimpleWorker(jobStore)
	jobScheduler := scheduler.NewSimpleScheduler(jobStore, jobWorker, schedulerIntervalSeconds)

	jobScheduler.Start()
	defer jobScheduler.Stop()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	setupRoutes(router)

	go func() {
		if err := router.Run(serverPort); err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func setupRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API handlers would be set up here
}
