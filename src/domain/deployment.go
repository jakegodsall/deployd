package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusRunning   Status = "running"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
)

type Deployment struct {
	ID           uuid.UUID `json:"id"`
	Application  string    `json:"application"`
	Environment  string    `json:"environment"`
	Version      uint32    `json:"version"`
	Status       Status    `json:"status"`
	StartedAt    time.Time `json:"started_at"`
	FinishedAt  *time.Time `json:"finished_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewDeployment(
	application string,
	environment string,
	version     uint32,
) *Deployment {
	now := time.Now()

	return &Deployment{
		ID:          uuid.New(),
		Application: application,
		Environment: environment,
		Version:     version,
		Status:      StatusRunning,
		StartedAt:   now,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (d Deployment) String() string {
	return fmt.Sprintf(
		"%s %s v%d [%s]",
		d.Application,
		d.Environment,
		d.Version,
		d.Status,
	)
}
