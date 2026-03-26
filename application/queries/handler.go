package queries

import (
	"context"

	"hexagonal-go/domain/entities"
	"hexagonal-go/domain/ports/outbound"
	"hexagonal-go/domain/valueobjects"
)

// QueryHandler handles read operations (CQRS Queries)
type QueryHandler struct {
	repo outbound.Repository
}

// NewQueryHandler creates a new query handler
func NewQueryHandler(repo outbound.Repository) *QueryHandler {
	return &QueryHandler{repo: repo}
}

// GetByIDQuery retrieves an entity by ID
type GetByIDQuery struct {
	ID string
}

// Handle executes the get by ID query
func (h *QueryHandler) HandleGetByID(ctx context.Context, query GetByIDQuery) (*entities.Example, error) {
	return h.repo.FindByID(ctx, query.ID)
}

// ListQuery returns paginated entities
type ListQuery struct {
	Page     int
	PageSize int
}

// Handle executes the list query
func (h *QueryHandler) HandleList(ctx context.Context, query ListQuery) ([]*entities.Example, error) {
	page := query.Page
	pageSize := query.PageSize
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}
	pagination := valueobjects.Pagination{
		Page:     page,
		PageSize: pageSize,
	}
	return h.repo.List(ctx, pagination)
}
