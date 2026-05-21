package domain

import (
	"time"

	"github.com/google/uuid"
)

// Lesson -------------------------------------------------------------------------------------------------------------

type Lesson struct {
	LessonID  int       `db:"lesson_id" json:"lesson_id"`
	TeacherID uuid.UUID `db:"teacher_id" json:"teacher_id"`
	Date      time.Time `db:"date" json:"date"`
	GroupID   int       `db:"group_id" json:"group_id"`
	SubjectID int       `db:"subject_id" json:"subject_id"`
	Room      int       `db:"room" json:"room"`
}

func NewLesson(li int, ti uuid.UUID, d time.Time, gi int, si int, r int) Lesson {
	return Lesson{li, ti, d, gi, si, r}
}
func (l *Lesson) SetTeacher(nti uuid.UUID) {
	l.TeacherID = nti
}
func (l *Lesson) SetDate(nd time.Time) {
	l.Date = nd
}
func (l *Lesson) SetGroup(nsgi int) {
	l.GroupID = nsgi
}
func (l *Lesson) SetSubject(nsgi int) {
	l.SubjectID = nsgi
}
func (l *Lesson) SetRoom(nr int) {
	l.Room = nr
}
func (l Lesson) GetID() int { return l.LessonID }

// --------------------------------------------------------------------------------------------------------------------
