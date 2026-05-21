package repository

import (
	"IoT-SmartPens/server_part/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Subject Manager -----------------------------------------------------------------------------------------------------

type SubjectManager struct {
	db *pgxpool.Pool
}

func NewSubjectManager(db *pgxpool.Pool) (*SubjectManager, error) {
	if db == nil {
		return nil, errors.New("db pgxpool.Pool is nil. NewSubjectManager()")
	}
	return &SubjectManager{db: db}, nil
}

var subjectsTableConfig = tableConfig{
	table:       "subjects",
	id_name:     "subject_id",
	oth_id_name: "",
}

func (m *SubjectManager) Insert(ctx context.Context, subject domain.Subject) error {
	query := `
        INSERT INTO subjects (name)
        VALUES ($1)
    `
	_, err := m.db.Exec(ctx, query, subject.Name)
	return err
}

func (m *SubjectManager) GetByID(ctx context.Context, id int) (*domain.Subject, error) {
	query := `SELECT subject_id, name FROM subjects WHERE subject_id = $1`
	res := m.db.QueryRow(ctx, query, id)
	subject := &domain.Subject{}
	err := res.Scan(&subject.SubjectID, &subject.Name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown subject_id")
	}
	return subject, nil
}

func (m *SubjectManager) DeleteByID(ctx context.Context, id int) error {
	return deleteByIDFunctionalWL(ctx, m.db, id, subjectsTableConfig)
}

func (m *SubjectManager) ExistsByID(ctx context.Context, id int) (bool, error) {
	return existsByIDFunctionalWL(ctx, m.db, id, subjectsTableConfig)
}

func (m *SubjectManager) GetByName(ctx context.Context, name string) (*domain.Subject, error) {
	query := `SELECT subject_id, name FROM subjects WHERE name = $1`
	res := m.db.QueryRow(ctx, query, name)
	subject := &domain.Subject{}
	err := res.Scan(&subject.SubjectID, &subject.Name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown subject name")
	}
	return subject, nil
}

func (m *SubjectManager) GetAll(ctx context.Context) ([]domain.Subject, error) {
	query := `SELECT subject_id, name FROM subjects ORDER BY subject_id`
	rows, err := m.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []domain.Subject
	for rows.Next() {
		var subject domain.Subject
		err := rows.Scan(&subject.SubjectID, &subject.Name)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}

// --------------------------------------------------------------------------------------------------------------------
