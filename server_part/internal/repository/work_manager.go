package repository

import (
	"IoT-SmartPens/server_part/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Work Manager -------------------------------------------------------------------------------------------------------

type WorkManager struct {
	db *pgxpool.Pool
}

func NewWorkManager(db *pgxpool.Pool) (*WorkManager, error) {
	if db == nil {
		return nil, errors.New("db pgxpool.Pool is nil. NewWorkManager()")
	}
	return &WorkManager{db: db}, nil
}

// Used only for Exist and Delete
var worksTableConfig = tableConfig{
	table:       "works",
	id_name:     "work_id",
	oth_id_name: "",
}

func (m *WorkManager) Insert(ctx context.Context, work domain.Work) error {
	query := `
        INSERT INTO works (lesson_id, student_id, data, mark)
        VALUES ($1, $2, $3, $4)
    `
	_, err := m.db.Exec(ctx, query, work.LessonID, work.StudentID, work.Data, work.Mark)
	return err
}

func (m *WorkManager) GetByID(ctx context.Context, id int) (*domain.Work, error) {
	query := `SELECT work_id, lesson_id, student_id, data, mark FROM works WHERE work_id = $1`
	res := m.db.QueryRow(ctx, query, id)
	work := &domain.Work{}
	err := res.Scan(&work.WorkID, &work.LessonID, &work.StudentID, &work.Data, &work.Mark)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown work_id")
	}
	return work, nil
}
func (m *WorkManager) DeleteByID(ctx context.Context, id int) error {
	return deleteByIDFunctionalWL(ctx, m.db, id, worksTableConfig)
}
func (m *WorkManager) ExistsByID(ctx context.Context, id int) (bool, error) {
	return existsByIDFunctionalWL(ctx, m.db, id, worksTableConfig)
}

// --------------------------------------------------------------------------------------------------------------------
