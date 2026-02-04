package store

import "errors"

var (
	ErrJobAlreadyExists           = errors.New("job already exists")
	ErrJobNotFound                = errors.New("job not found")
	ErrInvalidJobInput            = errors.New("invalid job input")
	ErrInvalidJobStatusTransition = errors.New("invalid job status transition")
)
