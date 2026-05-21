package service

import (
	"IoT-SmartPens/server_part/internal/domain"
	"IoT-SmartPens/server_part/internal/repository"
	"context"
)

// Subject Service -----------------------------------------------------------------------------------------------------

type SubjectService struct {
	*BaseIntService[domain.Subject]
}

func NewSubjectService(repo *repository.SubjectManager) *SubjectService {
	return &SubjectService{
		BaseIntService: NewBaseIntService(repo),
	}
}

func (s *SubjectService) UpdateName(ctx context.Context, id int, name string) error {
	return s.UpdateField(ctx, id, func(sub *domain.Subject) {
		sub.Name = name
	})
}

// --------------------------------------------------------------------------------------------------------------------
