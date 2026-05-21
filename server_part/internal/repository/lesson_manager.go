package repository

import (
	"IoT-SmartPens/server_part/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Lesson Manager -----------------------------------------------------------------------------------------------------

type LessonManager struct {
	db *pgxpool.Pool
}

func NewLessonManager(db *pgxpool.Pool) (*LessonManager, error) {
	if db == nil {
		return nil, errors.New("db pgxpool.Pool is nil. NewLessonManager()")
	}
	return &LessonManager{db: db}, nil
}

// Used only for Exist and Delete
var lessonsTableConfig = tableConfig{
	table:       "lessons",
	id_name:     "lesson_id",
	oth_id_name: "",
}

func (m *LessonManager) Insert(ctx context.Context, less domain.Lesson) error {
	query := `
        INSERT INTO lessons (lesson_id, teacher_id, date, group_id, subject_id, room)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (lesson_id)
        DO UPDATE SET
            teacher_id = EXCLUDED.teacher_id,
            date = EXCLUDED.date,
            group_id = EXCLUDED.group_id,
            subject_id = EXCLUDED.subject_id
			room = EXCLUDED.room
    `
	_, err := m.db.Exec(ctx, query, less.LessonID, less.TeacherID, less.Date, less.GroupID, less.SubjectID, less.Room)
	return err
}

func (m *LessonManager) GetByID(ctx context.Context, id int) (*domain.Lesson, error) {
	query := `SELECT lesson_id, teacher_id, date, group_id, subject_id, room FROM lessons WHERE lesson_id = $1`
	res := m.db.QueryRow(ctx, query, id)
	less := &domain.Lesson{}
	err := res.Scan(&less.LessonID, &less.TeacherID, &less.Date, &less.GroupID, &less.SubjectID, &less.Room)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown lesson_id")
	}
	return less, nil
}
func (m *LessonManager) DeleteByID(ctx context.Context, id int) error {
	return deleteByIDFunctionalWL(ctx, m.db, id, lessonsTableConfig)
}
func (m *LessonManager) ExistsByID(ctx context.Context, id int) (bool, error) {
	return existsByIDFunctionalWL(ctx, m.db, id, lessonsTableConfig)
}

// --------------------------------------------------------------------------------------------------------------------
