package repository

import (
	"context"
	"smartPens/internal/domain"
)

type LessonRepository interface {
	Create(ctx context.Context, lesson *domain.Lesson) error
	GetByID(ctx context.Context, id int) (*domain.Lesson, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]domain.Lesson, error)
}
