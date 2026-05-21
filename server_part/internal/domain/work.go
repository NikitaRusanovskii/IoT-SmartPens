package domain

import "github.com/google/uuid"

// Work ---------------------------------------------------------------------------------------------------------------

type Coords struct {
	Xcoord   int16
	Ycoord   int16
	Pressure int16
}
type Work struct {
	WorkID    int       `db:"work_id"`
	LessonID  int       `db:"lesson_id"`
	StudentID uuid.UUID `db:"student_id" json:"student_id"`
	Data      []Coords  `db:"data" json:"data"`
	Mark      int       `db:"mark"`
}

func NewWork(wi int, li int, ai uuid.UUID, d []Coords, m int) Work {
	return Work{wi, li, ai, d, m}
}
func (w *Work) SetLesson(nl int) {
	w.LessonID = nl
}
func (w *Work) SetMark(nm int) {
	w.Mark = nm
}
func (w Work) GetID() int { return w.WorkID }

// --------------------------------------------------------------------------------------------------------------------
