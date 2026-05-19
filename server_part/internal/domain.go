package internal

import (
	"time"

	"github.com/google/uuid"
)

type Role string

// Может быть добавлена роль ученика в будущем для доступа с ограниченными полномочиями
const (
	TeacherRole Role = "teacher"
	//StudentRole Role = "student"
)

type User struct {
	UserID         uuid.UUID `db:"user_id"`
	Role           Role      `db:"role"`
	IsOnline       bool      `db:"is_online"`
	AddrPort       string    `db:"addr_port"`
	ConnectionTime time.Time `db:"connection_time"`
}

func NewUser(Id uuid.UUID, Rl Role, Io bool, Ap string, Ct time.Time) User {
	return User{Id, Rl, Io, Ap, Ct}
}
func (u *User) SetRole(nr Role) {
	u.Role = nr
}

type Work struct {
	WorkID   int       `db:"work_id"`
	LessonID int       `db:"lesson_id"`
	AuthorID uuid.UUID `db:"user_id"`
	ImageURL string    `db:"image_url"`
}

func NewWork(wi int, li int, ai uuid.UUID, iu string) Work {
	return Work{wi, li, ai, iu}
}
func (w *Work) SetLesson(nl int) {
	w.LessonID = nl
}
func (w *Work) SetAuthor(na uuid.UUID) {
	w.AuthorID = na
}
func (w *Work) SetImageURL(niu string) {
	w.ImageURL = niu
}

type Lesson struct {
	LessonID  int       `db:"lesson_id"`
	TeacherID uuid.UUID `db:"user_id"`
}

func NewLesson(li int, ti uuid.UUID) Lesson {
	return Lesson{li, ti}
}
func (l *Lesson) SetTeacher(nti uuid.UUID) {
	l.TeacherID = nti
}
