package repository

import (
	"context"
	"errors"

	"smartPens/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrLessonNotFound = errors.New("lesson not found")
)

type LessonRepository struct {
	pool *pgxpool.Pool
}

func NewLessonRepository(pool *pgxpool.Pool) *LessonRepository {
	return &LessonRepository{
		pool: pool,
	}
}

func (r *LessonRepository) Create(
	ctx context.Context,
	lesson *domain.Lesson,
) error {

	query := `
		INSERT INTO lessons (
			teacher_id,
			group_id,
			subject_id,
			room,
			date
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	return r.pool.QueryRow(
		ctx,
		query,
		lesson.TeacherID,
		lesson.GroupID,
		lesson.SubjectID,
		lesson.Room,
		lesson.Date,
	).Scan(&lesson.ID)
}

func (r *LessonRepository) GetByID(
	ctx context.Context,
	id int,
) (*domain.Lesson, error) {

	query := `
		SELECT
			id,
			teacher_id,
			group_id,
			subject_id,
			room,
			date
		FROM lessons
		WHERE id = $1
	`

	var lesson domain.Lesson

	err := r.pool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&lesson.ID,
		&lesson.TeacherID,
		&lesson.GroupID,
		&lesson.SubjectID,
		&lesson.Room,
		&lesson.Date,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrLessonNotFound
		}

		return nil, err
	}

	return &lesson, nil
}

func (r *LessonRepository) Delete(
	ctx context.Context,
	id int,
) error {

	query := `
		DELETE FROM lessons
		WHERE id = $1
	`

	cmdTag, err := r.pool.Exec(
		ctx,
		query,
		id,
	)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrLessonNotFound
	}

	return nil
}

func (r *LessonRepository) List(
	ctx context.Context,
) ([]domain.Lesson, error) {

	query := `
		SELECT
			id,
			teacher_id,
			group_id,
			subject_id,
			room,
			date
		FROM lessons
		ORDER BY date
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var lessons []domain.Lesson

	for rows.Next() {
		var lesson domain.Lesson

		err := rows.Scan(
			&lesson.ID,
			&lesson.TeacherID,
			&lesson.GroupID,
			&lesson.SubjectID,
			&lesson.Room,
			&lesson.Date,
		)

		if err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lessons, nil
}
