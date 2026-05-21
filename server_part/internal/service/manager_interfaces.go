package service

import (
	"context"

	"github.com/google/uuid"
)

// Managers Interfaces ------------------------------------------------------------------------------------------------

type UUIDIdentifiable interface {
	GetID() uuid.UUID
}
type IntIdentifiable interface {
	GetID() int
}

type UUIDManager[T any] interface {
	Insert(ctx context.Context, entity T) error
	GetByID(ctx context.Context, id uuid.UUID) (*T, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	ExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
}
type IntManager[T any] interface {
	Insert(ctx context.Context, entity T) error
	GetByID(ctx context.Context, id int) (*T, error)
	DeleteByID(ctx context.Context, id int) error
	ExistsByID(ctx context.Context, id int) (bool, error)
}

// --------------------------------------------------------------------------------------------------------------------
