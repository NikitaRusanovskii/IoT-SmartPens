package repository

import (
	"IoT-SmartPens/server_part/internal/domain"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Student Manager ----------------------------------------------------------------------------------------------------

type StudentManager struct {
	db *pgxpool.Pool
}

func NewStudentManager(db *pgxpool.Pool) (*StudentManager, error) {
	if db == nil {
		return nil, errors.New("db pgxpool.Pool is nil. NewStudentManager()")
	}
	return &StudentManager{db: db}, nil
}

var studentsTableConfig = tableConfig{
	table:       "students",
	id_name:     "student_id",
	oth_id_name: "group_id",
}

func (m *StudentManager) Insert(ctx context.Context, stud domain.Student) error {
	return insertFunctional(ctx, m.db, stud.StudentID, stud.FIO, stud.GroupID, studentsTableConfig)
}
func (m *StudentManager) GetByID(ctx context.Context, id uuid.UUID) (*domain.Student, error) {
	res := getByIDFunctional(ctx, m.db, id, studentsTableConfig)
	stud := &domain.Student{}
	err := res.Scan(&stud.StudentID, &stud.FIO.FirstName, &stud.FIO.MiddleName, &stud.FIO.LastName, &stud.GroupID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown student_id")
	}
	return stud, nil
}
func (m *StudentManager) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return deleteByIDFunctionalST(ctx, m.db, id, studentsTableConfig)
}
func (m *StudentManager) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return existsByIDFunctionalST(ctx, m.db, id, studentsTableConfig)
}

// --------------------------------------------------------------------------------------------------------------------
