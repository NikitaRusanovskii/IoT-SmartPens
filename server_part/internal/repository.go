package internal

import (
	"context"
	"errors"

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

// Student Manager ----------------------------------------------------------------------------------------------------

type StudentManager struct {
	db *pgxpool.Pool
}

func NewStudentManager(db *pgxpool.Pool) (*StudentManager, error) {
	if db == nil {
		return nil, errors.New("db pgxpool.Pool is nil. NewUserManager()")
	}
	return &StudentManager{db: db}, nil
}
func (um *StudentManager) Insert(ctx context.Context, stud Student) error {
	query := `
		INSERT INTO students (student_id, fname, mname, lname, group_id)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT
			DO UPDATE SET
			fname = EXCLUDED.fname,
			mname = EXCLUDED.mname,
			lname = EXCLUDED.lname,
			group_id = EXCLUDED.group_id
	`
	_, err := um.db.Exec(ctx, query, stud.StudentID, stud.FirstName, stud.MiddleName, stud.LastName, stud.GroupID)
	return err
}

// type Student struct {
// 	StudentID  uuid.UUID `db:"student_id"`
// 	FirstName  string    `db:"fname"`
// 	MiddleName string    `db:"mname"`
// 	LastName   string    `db:"lname"`
// 	GroupID    int       `db:"group_id"`
// }

func (um *StudentManager) GetByID(ctx context.Context, id uuid.UUID) (*Student, error) {
	query := `
	SELECT student_id, fname, mname, lname, group_id
	FROM students
	WHERE id = $1
	`
	res := um.db.QueryRow(ctx, query, id)
	stud := &Student{}
	err := res.Scan(&stud.StudentID, &stud.FirstName, &stud.MiddleName, &stud.LastName, &stud.GroupID)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown user id")
	}
	return stud, nil
}

// --------------------------------------------------------------------------------------------------------------------
