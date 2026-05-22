package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// UUIDService --------------------------------------------------------------------------------------------------------

type BaseUUIDService[T UUIDIdentifiable] struct {
	repo UUIDManager[T]
}

func NewBaseUUIDService[T UUIDIdentifiable](repo UUIDManager[T]) *BaseUUIDService[T] {
	return &BaseUUIDService[T]{repo: repo}
}

func (s *BaseUUIDService[T]) Register(ctx context.Context, entity T) error {
	isExists, err := s.repo.ExistsByID(ctx, entity.GetID())
	if err != nil {
		return err
	}
	if isExists {
		return errors.New("already registered")
	}
	return s.repo.Insert(ctx, entity)
}

func (s *BaseUUIDService[T]) UpdateField(ctx context.Context, id uuid.UUID, updater func(*T)) error {
	isExists, err := s.repo.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !isExists {
		return errors.New("does not exist")
	}
	entity, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	updater(entity)
	return s.repo.Insert(ctx, *entity)
}

func (s *BaseUUIDService[T]) Remove(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteByID(ctx, id)
}

func (s *BaseUUIDService[T]) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.repo.ExistsByID(ctx, id)
}

func (s *BaseUUIDService[T]) GetByID(ctx context.Context, id uuid.UUID) (*T, error) {
	return s.repo.GetByID(ctx, id)
}

// --------------------------------------------------------------------------------------------------------------------

// IntService ---------------------------------------------------------------------------------------------------------

type BaseIntService[T IntIdentifiable] struct {
	repo IntManager[T]
}

func NewBaseIntService[T IntIdentifiable](repo IntManager[T]) *BaseIntService[T] {
	return &BaseIntService[T]{repo: repo}
}

func (s *BaseIntService[T]) Register(ctx context.Context, entity T) error {
	isExists, err := s.repo.ExistsByID(ctx, entity.GetID())
	if err != nil {
		return err
	}
	if isExists {
		return errors.New("already registered")
	}
	return s.repo.Insert(ctx, entity)
}

func (s *BaseIntService[T]) UpdateField(ctx context.Context, id int, updater func(*T)) error {
	isExists, err := s.repo.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !isExists {
		return errors.New("does not exist")
	}
	entity, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	updater(entity)
	return s.repo.Insert(ctx, *entity)
}

func (s *BaseIntService[T]) Remove(ctx context.Context, id int) error {
	return s.repo.DeleteByID(ctx, id)
}

func (s *BaseIntService[T]) ExistsByID(ctx context.Context, id int) (bool, error) {
	return s.repo.ExistsByID(ctx, id)
}

func (s *BaseIntService[T]) GetByID(ctx context.Context, id int) (*T, error) {
	return s.repo.GetByID(ctx, id)
}

// --------------------------------------------------------------------------------------------------------------------
