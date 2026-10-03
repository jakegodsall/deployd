package memory

import (
	"errors"
	"jakegodsall/deployd/src/domain"
)

var ErrNotFound = errors.New("deployment not found")

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

func (r *DeploymentRepository) FindAllByApplication(application string) ([]*domain.Deployment, error) {
	res := []*domain.Deployment{}

	for _, d := range r.deployments {
		if d.Application == application {
			res = append(res, d)
		}
	}

	return res, nil
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

func (r *DeploymentRepository) FindByApplicationVersion(application string, version uint32) (*domain.Deployment, error) {
	for _, d := range r.deployments {
		if d.Application == application && d.Version == version {
			return d, nil
		}
	}

	return nil, ErrNotFound
}
