package service

import (
	"context"
	"smartPens/internal/domain"
	"smartPens/internal/repository"
	"time"
)

type LearningService struct {
	workDB   *repository.WorkRepository
	lessonDB *repository.LessonRepository
}

func (l *LearningService) Delete(ctx context.Context, lessonID int, studentID int) error {
	err := l.workDB.Delete(ctx, lessonID, studentID)
	return err
}

func (l *LearningService) StartLesson(ctx context.Context,
	lessonID int, teacherID int, groupID int, subjectID int, room int, date time.Time) error {
	lesson, err := domain.NewLesson(lessonID, teacherID, groupID, subjectID, room, date)
	if err != nil {
		return err
	}

	err = l.lessonDB.Create(ctx, lesson)
	if err != nil {
		return err
	}
	return nil
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
