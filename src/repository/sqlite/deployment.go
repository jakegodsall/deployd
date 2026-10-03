package sqlite

import (
	"database/sql"
	"jakegodsall/deployd/src/domain"
)

type DeploymentRepository struct {
	db *sql.DB
}

func NewDeploymentRepository(db *sql.DB) *DeploymentRepository {
	return &DeploymentRepository{
		db: db,
	}
}

func (r *DeploymentRepository) FindAll() ([]*domain.Deployment, error) {
	rows, err := r.db.Query(`
		SELECT application, environment, version, status, started_at, finished_at
		FROM deployments
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	deployments := []*domain.Deployment{}

	for rows.Next() {
		deployment := &domain.Deployment{}

		if err := rows.Scan(
			&deployment.Application,
			&deployment.Environment,
			&deployment.Version,
			&deployment.Status,
			&deployment.StartedAt,
			&deployment.FinishedAt,
		); err != nil {
			return nil, err
		}

		deployments = append(deployments, deployment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deployments, nil
}
