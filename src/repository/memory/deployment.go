package memory

import "jakegodsall/deployd/src/domain"

type DeploymentRepository struct {
	deployments []*domain.Deployment
}

func NewDeploymentRepository() (*DeploymentRepository) {
	return &DeploymentRepository{
		deployments: []*domain.Deployment{
		},
	}
}

func (r *DeploymentRepository) FindAll() ([]*domain.Deployment, error) {
	return r.deployments, nil
}
