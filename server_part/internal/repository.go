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

// Student & Teacher Managers General Functions -----------------------------------------------------------------------

type tableConfig struct {
	table       string
	id_name     string
	oth_id_name string
}

// TO DO: protection from SQL injection (use pgx.Identifier{...}.Sanitize())

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
func deleteByIDFunctionalST(ctx context.Context, db *pgxpool.Pool, id uuid.UUID, tc tableConfig) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s = $1`, tc.table, tc.id_name)
	_, err := db.Exec(ctx, query, id)
	return err
}
func existsByIDFunctionalST(ctx context.Context, db *pgxpool.Pool, id uuid.UUID, tc tableConfig) (bool, error) {
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

// Work & Lesson Managers General Functions ---------------------------------------------------------------------------

func deleteByIDFunctionalWL(ctx context.Context, db *pgxpool.Pool, id int, tc tableConfig) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s = $1`, tc.table, tc.id_name)
	_, err := db.Exec(ctx, query, id)
	return err
}
func existsByIDFunctionalWL(ctx context.Context, db *pgxpool.Pool, id int, tc tableConfig) (bool, error) {
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
	return deleteByIDFunctionalST(ctx, m.db, id, studentsTableConfig)
}
func (m *StudentManager) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return existsByIDFunctionalST(ctx, m.db, id, studentsTableConfig)
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
	oth_id_name: "subject_id",
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
	return deleteByIDFunctionalST(ctx, m.db, id, teachersTableConfig)
}
func (m *TeacherManager) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return existsByIDFunctionalST(ctx, m.db, id, teachersTableConfig)
}

// --------------------------------------------------------------------------------------------------------------------

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

func (m *WorkManager) Insert(ctx context.Context, work Work) error {
	query := `
        INSERT INTO works (work_id, lesson_id, student_id, data, mark)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (work_id)
        DO UPDATE SET
            lesson_id = EXCLUDED.lesson_id,
            student_id = EXCLUDED.student_id,
            data = EXCLUDED.data,
            mark = EXCLUDED.mark
    `
	_, err := m.db.Exec(ctx, query, work.WorkID, work.LessonID, work.StudentID, work.Data, work.Mark)
	return err
}

func (m *WorkManager) GetByID(ctx context.Context, id int) (*Work, error) {
	query := `SELECT work_id, lesson_id, student_id, data, mark FROM works WHERE work_id = $1`
	res := m.db.QueryRow(ctx, query, id)
	work := &Work{}
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

func (m *LessonManager) Insert(ctx context.Context, less Lesson) error {
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

func (m *LessonManager) GetByID(ctx context.Context, id int) (*Lesson, error) {
	query := `SELECT lesson_id, teacher_id, date, group_id, subject_id, room FROM lessons WHERE lesson_id = $1`
	res := m.db.QueryRow(ctx, query, id)
	less := &Lesson{}
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
