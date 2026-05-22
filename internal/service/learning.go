package service

import (
	"context"
	"smartPens/internal/domain"
	"smartPens/internal/repository"
	"time"

	"github.com/google/uuid"
)

type LearningService struct {
	workDB   *repository.WorkRepository
	lessonDB *repository.LessonRepository
}

func NewLearningService(workRepo *repository.WorkRepository, lessonRepo *repository.LessonRepository) *LearningService {
	return &LearningService{
		workDB:   workRepo,
		lessonDB: lessonRepo,
	}
}
func (l *LearningService) Delete(ctx context.Context, lessonID int, studentID uuid.UUID) error {
	err := l.workDB.Delete(ctx, lessonID, studentID)
	return err
}

func (l *LearningService) StartLesson(ctx context.Context,
	teacherID uuid.UUID, groupID int, subjectID int, room int, date time.Time) (int, error) {
	lesson, err := domain.NewLesson(teacherID, groupID, subjectID, room, date)
	if err != nil {
		return 0, err
	}

	lessonID, err := l.lessonDB.Create(ctx, lesson)
	if err != nil {
		return 0, err
	}
	return lessonID, nil
}

func (l *LearningService) CompleteLesson(ctx context.Context, lessonID int, studentCnt int) error {
	works, err := GenerateDataFromStudentsPen(lessonID, studentCnt)
	if err != nil {
		return err
	}
	for _, w := range works {
		err := l.workDB.Create(ctx, &w)
		if err != nil {
			return err
		}
	}
	return nil
}
