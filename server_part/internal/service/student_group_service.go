package service

import (
	"IoT-SmartPens/server_part/internal/domain"
	"IoT-SmartPens/server_part/internal/repository"
	"context"
)

// StudentGroup Service ------------------------------------------------------------------------------------------------

type StudentGroupService struct {
	*BaseIntService[domain.StudentGroup]
}

func NewStudentGroupService(repo *repository.StudentGroupManager) *StudentGroupService {
	return &StudentGroupService{
		BaseIntService: NewBaseIntService(repo),
	}
}

func (s *StudentGroupService) UpdateName(ctx context.Context, id int, name string) error {
	return s.UpdateField(ctx, id, func(sg *domain.StudentGroup) {
		sg.Name = name
	})
}

// --------------------------------------------------------------------------------------------------------------------
