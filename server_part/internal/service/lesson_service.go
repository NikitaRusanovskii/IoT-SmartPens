package service

import (
	"IoT-SmartPens/server_part/internal/domain"
	"IoT-SmartPens/server_part/internal/repository"
	"context"
)

// Lesson Service -----------------------------------------------------------------------------------------------------

type LessonService struct {
	*BaseIntService[domain.Lesson]
}

func NewLessonService(repo *repository.LessonManager) *LessonService {
	return &LessonService{
		BaseIntService: NewBaseIntService(repo),
	}
}

func (s *LessonService) UpdateRoom(ctx context.Context, id int, room int) error {
	return s.UpdateField(ctx, id, func(l *domain.Lesson) {
		l.Room = room
	})
}

// --------------------------------------------------------------------------------------------------------------------
