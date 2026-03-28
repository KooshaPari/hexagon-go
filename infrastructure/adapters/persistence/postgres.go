package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"hexagonal-go/domain/entities"
	"hexagonal-go/domain/errors"
	"hexagonal-go/domain/ports/outbound"
	"hexagonal-go/domain/valueobjects"
)

// PostgresRepository is a PostgreSQL implementation of the Repository port
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository creates a new PostgresRepository
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Compile-time interface check
var _ outbound.Repository = (*PostgresRepository)(nil)

// Save persists an entity to PostgreSQL
func (r *PostgresRepository) Save(ctx context.Context, entity *entities.Example) error {
	query := `
		INSERT INTO entities (id, created_at, updated_at, name, description, active)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET
			updated_at = EXCLUDED.updated_at,
			name = EXCLUDED.name,
			description = EXCLUDED.description,
			active = EXCLUDED.active
	`
	_, err := r.db.ExecContext(ctx, query, entity.ID, entity.CreatedAt, entity.UpdatedAt, entity.Name, entity.Description, entity.Active)
	return err
}

// FindByID retrieves an entity by ID
func (r *PostgresRepository) FindByID(ctx context.Context, id string) (*entities.Example, error) {
	query := `SELECT id, created_at, updated_at, name, description, active FROM entities WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var entity entities.Example
	err := row.Scan(&entity.ID, &entity.CreatedAt, &entity.UpdatedAt, &entity.Name, &entity.Description, &entity.Active)
	if err == sql.ErrNoRows {
		return nil, errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete removes an entity
func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM entities WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.ErrNotFound
	}
	return nil
}

// List returns all entities with pagination
func (r *PostgresRepository) List(ctx context.Context, pagination valueobjects.Pagination) (result []*entities.Example, err error) {
	query := `
		SELECT id, created_at, updated_at, name, description, active
		FROM entities
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, pagination.Limit(), pagination.Offset())
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close rows: %w", closeErr)
		}
	}()

	result = make([]*entities.Example, 0)
	for rows.Next() {
		var entity entities.Example
		if err := rows.Scan(&entity.ID, &entity.CreatedAt, &entity.UpdatedAt, &entity.Name, &entity.Description, &entity.Active); err != nil {
			return nil, err
		}
		result = append(result, &entity)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// InitSchema creates the database schema
func (r *PostgresRepository) InitSchema(ctx context.Context) error {
	schema := `
		CREATE TABLE IF NOT EXISTS entities (
			id UUID PRIMARY KEY,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_entities_created_at ON entities(created_at);
	`
	_, err := r.db.ExecContext(ctx, schema)
	return err
}

// Close closes the database connection
func (r *PostgresRepository) Close() error {
	return r.db.Close()
}

// Wrap-Over: This adapter wraps a third-party SQL driver behind our interface
// This allows swapping the underlying database without changing domain code
type PostgresAdapter struct {
	*PostgresRepository
}

func NewPostgresAdapter(connectionString string) (*PostgresAdapter, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return &PostgresAdapter{
		PostgresRepository: NewPostgresRepository(db),
	}, nil
}
