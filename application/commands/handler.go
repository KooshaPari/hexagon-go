package commands

import (
	"context"

	"hexagonal-go/domain/entities"
	"hexagonal-go/domain/ports/outbound"
)

// CommandHandler handles write operations (CQRS Commands)
type CommandHandler struct {
	repo outbound.Repository
}

// NewCommandHandler creates a new command handler
func NewCommandHandler(repo outbound.Repository) *CommandHandler {
	return &CommandHandler{repo: repo}
}

// CreateCommand creates a new entity
type CreateCommand struct {
	Name        string
	Description string
}

// Handle executes the create command
func (h *CommandHandler) Handle(ctx context.Context, cmd CreateCommand) (*entities.Example, error) {
	entity := entities.NewExample(cmd.Name, cmd.Description)
	if err := entity.Validate(); err != nil {
		return nil, err
	}
	if err := h.repo.Save(ctx, entity); err != nil {
		return nil, err
	}
	return entity, nil
}

// UpdateCommand updates an existing entity
type UpdateCommand struct {
	ID          string
	Name        string
	Description string
}

// HandleUpdate executes the update command
func (h *CommandHandler) HandleUpdate(ctx context.Context, cmd UpdateCommand) (*entities.Example, error) {
	entity, err := h.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	entity.Name = cmd.Name
	entity.Description = cmd.Description
	entity.Touch()

	if err := h.repo.Save(ctx, entity); err != nil {
		return nil, err
	}
	return entity, nil
}

// DeleteCommand removes an entity
type DeleteCommand struct {
	ID string
}

// HandleDelete executes the delete command
func (h *CommandHandler) HandleDelete(ctx context.Context, cmd DeleteCommand) error {
	return h.repo.Delete(ctx, cmd.ID)
}
