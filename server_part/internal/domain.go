package internal

import (
	"time"

	"github.com/google/uuid"
)

// FIO struct ---------------------------------------------------------------------------------------------------------

type Fio struct {
	FirstName  string `db:"fname"`
	MiddleName string `db:"mname"`
	LastName   string `db:"lname"`
}

// --------------------------------------------------------------------------------------------------------------------

// Student ------------------------------------------------------------------------------------------------------------

type Student struct {
	StudentID uuid.UUID `db:"student_id"`
	FIO       Fio
	GroupID   int `db:"group_id"`
}

func NewStudent(Id uuid.UUID, fn string, mn string, ln string, gi int) Student {
	return Student{Id, Fio{fn, mn, ln}, gi}
}
func (s *Student) SetGroup(ng int) {
	s.GroupID = ng
}

// --------------------------------------------------------------------------------------------------------------------

// Teacher ------------------------------------------------------------------------------------------------------------

type Teacher struct {
	TeacherID uuid.UUID `db:"teacher_id"`
	FIO       Fio
	SubjectID int `db:"subject_id"`
}

func NewTeacher(Id uuid.UUID, fn string, mn string, ln string, si int) Teacher {
	return Teacher{Id, Fio{fn, mn, ln}, si}
}
func (t *Teacher) SetSubject(ns int) {
	t.SubjectID = ns
}

// --------------------------------------------------------------------------------------------------------------------

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

// --------------------------------------------------------------------------------------------------------------------

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

// --------------------------------------------------------------------------------------------------------------------
