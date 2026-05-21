package domain

import (
	"errors"
	"time"
)

type Lesson struct {
	ID        int       `json:"id" db:"id"`
	TeacherID int       `json:"teacher_id" db:"teacher_id"`
	GroupID   int       `json:"group_id" db:"group_id"`
	SubjectID int       `json:"subject_id" db:"subject_id"`
	Room      int       `json:"room" db:"room"`
	Date      time.Time `json:"date" db:"date"`
}

func NewLesson(
	teacherID int,
	groupID int,
	subjectID int,
	room int,
	date time.Time,
) (*Lesson, error) {

	if teacherID <= 0 {
		return nil, errors.New("invalid teacher id")
	}

	if groupID <= 0 {
		return nil, errors.New("invalid group id")
	}

	if subjectID <= 0 {
		return nil, errors.New("invalid subject id")
	}

	if room <= 0 {
		return nil, errors.New("invalid room")
	}

	if date.IsZero() {
		return nil, errors.New("date is required")
	}

	return &Lesson{
		TeacherID: teacherID,
		GroupID:   groupID,
		SubjectID: subjectID,
		Room:      room,
		Date:      date,
	}, nil
}
