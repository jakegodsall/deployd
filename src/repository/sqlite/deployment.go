package sqlite

import (
	"database/sql"
	"errors"
	"jakegodsall/deployd/src/domain"
	"jakegodsall/deployd/src/repository"
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
		SELECT id, application, environment, version, status, started_at, finished_at, created_at, updated_at
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
			&deployment.ID,
			&deployment.Application,
			&deployment.Environment,
			&deployment.Version,
			&deployment.Status,
			&deployment.StartedAt,
			&deployment.FinishedAt,
			&deployment.CreatedAt,
			&deployment.UpdatedAt,
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

func (r *DeploymentRepository) FindAllByApplication(application string) ([]*domain.Deployment, error) {
	rows, err := r.db.Query(`
		SELECT id, application, environment, version, status, started_at, finished_at, created_at, updated_at
		FROM deployments
		WHERE application = ?
	`, application)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	deployments := []*domain.Deployment{}

	for rows.Next() {
		deployment := &domain.Deployment{}

		if err := rows.Scan(
			&deployment.ID,
			&deployment.Application,
			&deployment.Environment,
			&deployment.Version,
			&deployment.Status,
			&deployment.StartedAt,
			&deployment.FinishedAt,
			&deployment.CreatedAt,
			&deployment.UpdatedAt,
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

func (r *DeploymentRepository) FindAllByStatus(status domain.Status) ([]*domain.Deployment, error) {
	rows, err := r.db.Query(`
        SELECT id, application, environment, version, status, started_at, finished_at, created_at, updated_at
        FROM deployments
        WHERE status = ?
	`, status)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	deployments := []*domain.Deployment{}

	for rows.Next() {
		deployment := &domain.Deployment{}

		if err := rows.Scan(
			&deployment.ID,
			&deployment.Application,
			&deployment.Environment,
			&deployment.Version,
			&deployment.Status,
			&deployment.StartedAt,
			&deployment.FinishedAt,
			&deployment.CreatedAt,
			&deployment.UpdatedAt,
		); err != nil {
			return nil, err
		}

		deployments = append(deployments, deployment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deployments, err
}

func (r *DeploymentRepository) FindByApplicationVersion(application string, version uint32) (*domain.Deployment, error) {
	row := r.db.QueryRow(`
		SELECT id, application, environment, version, status, started_at, finished_at, created_at, updated_at
		FROM deployments
		WHERE application = ? AND version = ?
	`, application, version)

	deployment := &domain.Deployment{}

	if err := row.Scan(
		&deployment.ID,
		&deployment.Application,
		&deployment.Environment,
		&deployment.Version,
		&deployment.Status,
		&deployment.StartedAt,
		&deployment.FinishedAt,
		&deployment.CreatedAt,
		&deployment.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return deployment, nil
}
