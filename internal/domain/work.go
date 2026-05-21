package domain

import (
	"errors"

	"github.com/google/uuid"
)

type PointList [][3]int8

type Work struct {
	LessonID  int
	StudentID uuid.UUID
	Data      PointList
}

func NewWork(
	lessonID int,
	studentID uuid.UUID,
	data PointList,
) (*Work, error) {

	if lessonID <= 0 {
		return nil, errors.New("invalid lesson id")
	}

	return &Work{
		LessonID:  lessonID,
		StudentID: studentID,
		Data:      data,
	}, nil
}
