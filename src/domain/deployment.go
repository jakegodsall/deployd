package domain

import "time"

type Status string

const (
	StatusRunning   Status = "running"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
)

type Deployment struct {
	Application string
	Environment string
	Version     uint32
	Status      Status
	StartedAt   time.Time
	FinishedAt  *time.Time
}

func NewDeployment(
	application string,
	environment string,
	version     uint32,
) *Deployment {
	return &Deployment{
		Application: application,
		Environment: environment,
		Version:     version,
		Status:      StatusRunning,
		StartedAt:   time.Now(),
	}
}
