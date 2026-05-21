package internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Connection ---------------------------------------------------------------------------------------------------------

type ConnectionManager struct {
	db *pgxpool.Pool
}

func NewConnectionManager(db *pgxpool.Pool) *ConnectionManager {
	if db == nil {
		return &ConnectionManager{db: nil}
	}
	return &ConnectionManager{db: db}
}
func (c *ConnectionManager) Connect(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	c.db = pool
	return pool, nil
}
func (c *ConnectionManager) Disconnect() {
	c.db.Close()
}
func (c *ConnectionManager) GetPool() (*pgxpool.Pool, error) {
	if c.db == nil {
		return nil, errors.New("db pgxpool.Pool is nil. ConnectionManager:GetPool()")
	}
	return c.db, nil
}

// --------------------------------------------------------------------------------------------------------------------

// General Functions --------------------------------------------------------------------------------------------------
type tableConfig struct {
	table       string
	id_name     string
	oth_id_name string
}

// TO DO: protection from SQL injection

func insertFunctional(ctx context.Context, db *pgxpool.Pool,
	id uuid.UUID, fio Fio, other_id int, tc tableConfig) error {
	query := fmt.Sprintf(`
        INSERT INTO %s (%s, fname, mname, lname, %s)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (%s)
        DO UPDATE SET
            fname = EXCLUDED.fname,
            mname = EXCLUDED.mname,
            lname = EXCLUDED.lname,
            %s = EXCLUDED.%s
    `, tc.table, tc.id_name, tc.oth_id_name, tc.id_name, tc.oth_id_name, tc.oth_id_name)
	_, err := db.Exec(ctx, query, id, fio.FirstName, fio.MiddleName, fio.LastName, other_id)
	return err
}
func getByIDFunctional(ctx context.Context, db *pgxpool.Pool, id uuid.UUID, tc tableConfig) pgx.Row {
	query := fmt.Sprintf(`SELECT %s, fname, mname, lname, %s FROM %s WHERE %s = $1`, tc.id_name, tc.oth_id_name, tc.table, tc.id_name)
	return db.QueryRow(ctx, query, id)
}
func deleteByIDFunctional(ctx context.Context, db *pgxpool.Pool, id uuid.UUID, tc tableConfig) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s = $1`, tc.table, tc.id_name)
	_, err := db.Exec(ctx, query, id)
	return err
}
func existsByIDFunctional(ctx context.Context, db *pgxpool.Pool, id uuid.UUID, tc tableConfig) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1)`, tc.table, tc.id_name)
	var exists bool
	res := db.QueryRow(ctx, query, id)
	err := res.Scan(&exists)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, errors.New(err.Error())
	}
	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

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

func (m *StudentManager) Insert(ctx context.Context, stud Student) error {
	return insertFunctional(ctx, m.db, stud.StudentID, stud.FIO, stud.GroupID, studentsTableConfig)
}
func (m *StudentManager) GetByID(ctx context.Context, id uuid.UUID) (*Student, error) {
	res := getByIDFunctional(ctx, m.db, id, studentsTableConfig)
	stud := &Student{}
	err := res.Scan(&stud.StudentID, &stud.FIO.FirstName, &stud.FIO.MiddleName, &stud.FIO.LastName, &stud.GroupID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown student_id")
	}
	return stud, nil
}
func (m *StudentManager) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return deleteByIDFunctional(ctx, m.db, id, studentsTableConfig)
}
func (m *StudentManager) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return existsByIDFunctional(ctx, m.db, id, studentsTableConfig)
}

// --------------------------------------------------------------------------------------------------------------------

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
	oth_id_name: "department_id",
}

func (m *TeacherManager) Insert(ctx context.Context, teac Teacher) error {
	return insertFunctional(ctx, m.db, teac.TeacherID, teac.FIO, teac.SubjectID, teachersTableConfig)
}
func (m *TeacherManager) GetByID(ctx context.Context, id uuid.UUID) (*Teacher, error) {
	res := getByIDFunctional(ctx, m.db, id, teachersTableConfig)
	teac := &Teacher{}
	err := res.Scan(&teac.TeacherID, &teac.FIO.FirstName, &teac.FIO.MiddleName, &teac.FIO.LastName, &teac.SubjectID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown teacher_id")
	}
	return teac, nil
}
func (m *TeacherManager) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return deleteByIDFunctional(ctx, m.db, id, teachersTableConfig)
}
func (m *TeacherManager) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return existsByIDFunctional(ctx, m.db, id, teachersTableConfig)
}

// --------------------------------------------------------------------------------------------------------------------
