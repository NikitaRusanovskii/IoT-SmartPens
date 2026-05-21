package postgres

import (
	"context"
	"errors"

	"smartPens/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrWorkNotFound = errors.New("work not found")
)

type WorkRepository struct {
	pool *pgxpool.Pool
}

func NewWorkRepository(pool *pgxpool.Pool) *WorkRepository {
	return &WorkRepository{
		pool: pool,
	}
}

func (r *WorkRepository) Create(
	ctx context.Context,
	work *domain.Work,
) error {

	query := `
		INSERT INTO works (
			lesson_id,
			student_id,
			data
		)
		VALUES ($1, $2, $3)
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		work.LessonID,
		work.StudentID,
		work.Data,
	)

	return err
}

func (r *WorkRepository) GetByLessonAndStudent(
	ctx context.Context,
	lessonID int,
	studentID int,
) (*domain.Work, error) {

	query := `
		SELECT
			lesson_id,
			student_id,
			data
		FROM works
		WHERE lesson_id = $1
		  AND student_id = $2
	`

	var work domain.Work

	err := r.pool.QueryRow(
		ctx,
		query,
		lessonID,
		studentID,
	).Scan(
		&work.LessonID,
		&work.StudentID,
		&work.Data,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkNotFound
		}

		return nil, err
	}

	return &work, nil
}

func (r *WorkRepository) Delete(
	ctx context.Context,
	lessonID int,
	studentID int,
) error {

	query := `
		DELETE FROM works
		WHERE lesson_id = $1
		  AND student_id = $2
	`

	cmdTag, err := r.pool.Exec(
		ctx,
		query,
		lessonID,
		studentID,
	)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrWorkNotFound
	}

	return nil
}
