package repository

import (
	"IoT-SmartPens/server_part/internal/domain"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Teacher Manager ----------------------------------------------------------------------------------------------------

type TeacherManager struct {
	db *pgxpool.Pool
}

func NewTeacherManager(db *pgxpool.Pool) (*TeacherManager, error) {
	if db == nil {
		return nil, errors.New("db pgxpool.Pool is nil. NewTeacherManager()")
	}
	return &TeacherManager{db: db}, nil
}

var teachersTableConfig = tableConfig{
	table:       "teachers",
	id_name:     "teacher_id",
	oth_id_name: "subject_id",
}

func (m *TeacherManager) Insert(ctx context.Context, teac domain.Teacher) error {
	return insertFunctional(ctx, m.db, teac.TeacherID, teac.FIO, teac.SubjectID, teachersTableConfig)
}
func (m *TeacherManager) GetByID(ctx context.Context, id uuid.UUID) (*domain.Teacher, error) {
	res := getByIDFunctional(ctx, m.db, id, teachersTableConfig)
	teac := &domain.Teacher{}
	err := res.Scan(&teac.TeacherID, &teac.FIO.FirstName, &teac.FIO.MiddleName, &teac.FIO.LastName, &teac.SubjectID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown teacher_id")
	}
	return teac, nil
}
func (m *TeacherManager) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return deleteByIDFunctionalST(ctx, m.db, id, teachersTableConfig)
}
func (m *TeacherManager) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return existsByIDFunctionalST(ctx, m.db, id, teachersTableConfig)
}

// --------------------------------------------------------------------------------------------------------------------
