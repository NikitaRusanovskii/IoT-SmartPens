package service

import (
	"IoT-SmartPens/server_part/internal/domain"
	"IoT-SmartPens/server_part/internal/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

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

// --------------------------------------------------------------------------------------------------------------------

// --------------------------------------------------------------------------------------------------------------------

type StudentService struct {
	*BaseUUIDService[domain.Student]
}

func NewStudentService(repo *repository.StudentManager) *StudentService {
	return &StudentService{
		BaseUUIDService: NewBaseUUIDService[domain.Student](repo),
	}
}

func (s *StudentService) UpdateGroup(ctx context.Context, id uuid.UUID, groupID int) error {
	return s.UpdateField(ctx, id, func(stud *domain.Student) {
		stud.GroupID = groupID
	})
}

type TeacherService struct {
	*BaseUUIDService[domain.Teacher]
}

func NewTeacherService(repo *repository.TeacherManager) *TeacherService {
	return &TeacherService{
		BaseUUIDService: NewBaseUUIDService[domain.Teacher](repo),
	}
}

func (s *TeacherService) UpdateSubject(ctx context.Context, id uuid.UUID, subjectID int) error {
	return s.UpdateField(ctx, id, func(teac *domain.Teacher) {
		teac.SubjectID = subjectID
	})
}

type WorkService struct {
	*BaseIntService[domain.Work]
}

func NewWorkService(repo *repository.WorkManager) *WorkService {
	return &WorkService{
		BaseIntService: NewBaseIntService[domain.Work](repo),
	}
}

func (s *WorkService) UpdateMark(ctx context.Context, id int, mark int) error {
	return s.UpdateField(ctx, id, func(w *domain.Work) {
		w.Mark = mark
	})
}

type LessonService struct {
	*BaseIntService[domain.Lesson]
}

func NewLessonService(repo *repository.LessonManager) *LessonService {
	return &LessonService{
		BaseIntService: NewBaseIntService[domain.Lesson](repo),
	}
}

func (s *LessonService) UpdateRoom(ctx context.Context, id int, room int) error {
	return s.UpdateField(ctx, id, func(l *domain.Lesson) {
		l.Room = room
	})
}
