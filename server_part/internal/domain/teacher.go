package domain

import "github.com/google/uuid"

// Teacher ------------------------------------------------------------------------------------------------------------

type Teacher struct {
	TeacherID uuid.UUID `db:"teacher_id" json:"teacher_id"`
	FIO       Fio       `json:"fio"`
	SubjectID int       `db:"subject_id" json:"subject_id"`
}

func NewTeacher(Id uuid.UUID, fn string, mn string, ln string, si int) Teacher {
	return Teacher{Id, Fio{fn, mn, ln}, si}
}
func (t *Teacher) SetSubject(ns int) {
	t.SubjectID = ns
}
func (t Teacher) GetID() uuid.UUID { return t.TeacherID }

// --------------------------------------------------------------------------------------------------------------------
