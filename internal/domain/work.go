package domain

import "errors"

type PointList [][3]int8

type Work struct {
	LessonID  int
	StudentID int
	Data      PointList
}

func NewWork(
	lessonID int,
	studentID int,
	data PointList,
) (*Work, error) {

	if lessonID <= 0 {
		return nil, errors.New("invalid lesson id")
	}

	if studentID <= 0 {
		return nil, errors.New("invalid student id")
	}

	return &Work{
		LessonID:  lessonID,
		StudentID: studentID,
		Data:      data,
	}, nil
}
