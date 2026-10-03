package repository

import (
	"errors"
	"jakegodsall/deployd/src/domain"
)

var ErrNotFound = errors.New("deployment not found")

type DeploymentRepository interface {
	FindAll() ([]*domain.Deployment, error)
	FindAllByApplication(application string) ([]*domain.Deployment, error)
	FindAllByStatus(status domain.Status) ([]*domain.Deployment, error)

	FindByApplicationVersion(application string, version uint32) (*domain.Deployment, error)
}
