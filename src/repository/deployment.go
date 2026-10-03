package repository

import "jakegodsall/deployd/src/domain"

type DeploymentRepository interface {
	FindAll() ([]*domain.Deployment, error)
}
