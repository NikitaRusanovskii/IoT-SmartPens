package service

import (
	"IoT-SmartPens/server_part/internal/domain"
	"IoT-SmartPens/server_part/internal/repository"
	"context"
)

// Work Service -------------------------------------------------------------------------------------------------------

type WorkService struct {
	*BaseIntService[domain.Work]
}

func NewWorkService(repo *repository.WorkManager) *WorkService {
	return &WorkService{
		BaseIntService: NewBaseIntService(repo),
	}
}

func (s *WorkService) UpdateMark(ctx context.Context, id int, mark int) error {
	return s.UpdateField(ctx, id, func(w *domain.Work) {
		w.Mark = mark
	})
}

// --------------------------------------------------------------------------------------------------------------------
