package memory

import "jakegodsall/deployd/src/domain"

type DeploymentRepository struct {
	deployments []*domain.Deployment
}

func NewDeploymentRepository() *DeploymentRepository {
	return &DeploymentRepository{
		deployments: []*domain.Deployment{
			domain.NewDeployment("my-app", "staging", 1),
		},
	}
}

func (r *DeploymentRepository) FindAll() ([]*domain.Deployment, error) {
	return r.deployments, nil
}

func (r *DeploymentRepository) FindAllByStatus(status domain.Status) ([]*domain.Deployment, error) {
	res := []*domain.Deployment{}

	for _, d := range r.deployments {
		if d.Status == status {
			res = append(res, d)
		}
	}

	return res, nil
}
