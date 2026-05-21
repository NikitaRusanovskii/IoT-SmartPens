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
		INSERT INTO lesson (
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
		FROM lesson
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
		DELETE FROM lesson
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
		FROM lesson
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

func (r *LessonRepository) CollectWorks(ctx context.Context, lessonID int) (*domain.WorksRequest, error) {
	query := `
		SELECT 
	    l.id, 
	    l.teacher_id, 
	    l.date, 
	    l.group_id, 
	    l.subject_id, 
	    l.room,
	    	COALESCE(
	    	    json_agg(
	    	        json_build_object(
	    	            'student_id', w.student_id,
	    	            'coords', w.data
	    	        )
	    	    ) FILTER (WHERE w.student_id IS NOT NULL), 
	    	    '[]'
	    	) as works
		FROM lesson l
		LEFT JOIN work w ON l.id = w.lesson_id
		WHERE l.id = $1
		GROUP BY l.id;
	`

	var req domain.WorksRequest
	err := r.pool.QueryRow(ctx, query, lessonID).Scan(
		&req.LessonID,
		&req.TeacherID,
		&req.Date,
		&req.GroupID,
		&req.SubjectID,
		&req.Room,
		&req.Works,
	)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
