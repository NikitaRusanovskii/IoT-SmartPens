package domain

import (
	"github.com/google/uuid"
)

// Student ------------------------------------------------------------------------------------------------------------

type Student struct {
	StudentID uuid.UUID `db:"student_id" json:"student_id"`
	FIO       Fio       `json:"fio"`
	GroupID   int       `db:"group_id" json:"group_id"`
}

func NewStudent(Id uuid.UUID, fn string, mn string, ln string, gi int) Student {
	return Student{Id, Fio{fn, mn, ln}, gi}
}
func (s *Student) SetGroup(ng int) {
	s.GroupID = ng
}
func (s Student) GetID() uuid.UUID { return s.StudentID }

// --------------------------------------------------------------------------------------------------------------------
