package domain

import "github.com/google/uuid"

// Work ---------------------------------------------------------------------------------------------------------------

type Coords struct {
	Xcoord   int16 `json:"Xcoord"`
	Ycoord   int16 `json:"Ycoord"`
	Pressure int16 `json:"Pressure"`
}

type Work struct {
	WorkID    int       `db:"work_id" json:"work_id"`
	LessonID  int       `db:"lesson_id" json:"lesson_id"`
	StudentID uuid.UUID `db:"student_id" json:"student_id"`
	Data      []Coords  `db:"data" json:"data"`
	Mark      int       `db:"mark" json:"mark"`
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
