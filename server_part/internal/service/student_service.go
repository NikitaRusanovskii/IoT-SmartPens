package service

import (
	"IoT-SmartPens/server_part/internal/domain"
	"IoT-SmartPens/server_part/internal/repository"
	"context"

	"github.com/google/uuid"
)

// Student Service ----------------------------------------------------------------------------------------------------

type StudentService struct {
	*BaseUUIDService[domain.Student]
}

func NewStudentService(repo *repository.StudentManager) *StudentService {
	return &StudentService{
		BaseUUIDService: NewBaseUUIDService(repo),
	}
}

func (s *StudentService) UpdateGroup(ctx context.Context, id uuid.UUID, groupID int) error {
	return s.UpdateField(ctx, id, func(stud *domain.Student) {
		stud.GroupID = groupID
	})
}

// --------------------------------------------------------------------------------------------------------------------
