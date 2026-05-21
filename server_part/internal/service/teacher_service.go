package service

import (
	"IoT-SmartPens/server_part/internal/domain"
	"IoT-SmartPens/server_part/internal/repository"
	"context"

	"github.com/google/uuid"
)

// Teacher Service ----------------------------------------------------------------------------------------------------

type TeacherService struct {
	*BaseUUIDService[domain.Teacher]
}

func NewTeacherService(repo *repository.TeacherManager) *TeacherService {
	return &TeacherService{
		BaseUUIDService: NewBaseUUIDService(repo),
	}
}

func (s *TeacherService) UpdateSubject(ctx context.Context, id uuid.UUID, subjectID int) error {
	return s.UpdateField(ctx, id, func(teac *domain.Teacher) {
		teac.SubjectID = subjectID
	})
}

// --------------------------------------------------------------------------------------------------------------------
