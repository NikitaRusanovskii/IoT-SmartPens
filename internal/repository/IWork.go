package repository

import (
	"context"
	"smartPens/internal/domain"
)

type WorkRepository interface {
	Create(ctx context.Context, work *domain.Work) error
	GetByLessonAndStudent(
		ctx context.Context,
		lessonID int,
		studentID int,
	) (*domain.Work, error)

	Delete(
		ctx context.Context,
		lessonID int,
		studentID int,
	) error
}
