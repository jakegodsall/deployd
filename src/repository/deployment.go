package repository

import "jakegodsall/deployd/src/domain"

type DeploymentRepository interface {
	FindAll() ([]*domain.Deployment, error)
	FindAllByStatus(status domain.Status) ([]*domain.Deployment, error)
}
