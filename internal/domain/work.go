package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Point struct {
	X        int8 `json:"x"`
	Y        int8 `json:"y"`
	Pressure int8 `json:"pressure"`
}

type Work struct {
	LessonID  int
	StudentID uuid.UUID
	Data      []Point
}

func NewWork(
	lessonID int,
	studentID uuid.UUID,
	data []Point,
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
