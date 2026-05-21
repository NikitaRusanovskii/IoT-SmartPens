package repository

import (
	"IoT-SmartPens/server_part/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// StudentGroup Manager ------------------------------------------------------------------------------------------------

type StudentGroupManager struct {
	db *pgxpool.Pool
}

func NewStudentGroupManager(db *pgxpool.Pool) (*StudentGroupManager, error) {
	if db == nil {
		return nil, errors.New("db pgxpool.Pool is nil. NewStudentGroupManager()")
	}
	return &StudentGroupManager{db: db}, nil
}

var studentGroupsTableConfig = tableConfig{
	table:       "student_groups",
	id_name:     "group_id",
	oth_id_name: "",
}

func (m *StudentGroupManager) Insert(ctx context.Context, group domain.StudentGroup) error {
	query := `
        INSERT INTO student_groups (group_id, name)
        VALUES ($1, $2)
        ON CONFLICT (group_id)
        DO UPDATE SET
            name = EXCLUDED.name
    `
	_, err := m.db.Exec(ctx, query, group.GroupID, group.Name)
	return err
}

func (m *StudentGroupManager) GetByID(ctx context.Context, id int) (*domain.StudentGroup, error) {
	query := `SELECT group_id, name FROM student_groups WHERE group_id = $1`
	res := m.db.QueryRow(ctx, query, id)
	group := &domain.StudentGroup{}
	err := res.Scan(&group.GroupID, &group.Name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown group_id")
	}
	return group, nil
}

func (m *StudentGroupManager) DeleteByID(ctx context.Context, id int) error {
	return deleteByIDFunctionalWL(ctx, m.db, id, studentGroupsTableConfig)
}

func (m *StudentGroupManager) ExistsByID(ctx context.Context, id int) (bool, error) {
	return existsByIDFunctionalWL(ctx, m.db, id, studentGroupsTableConfig)
}

func (m *StudentGroupManager) GetByName(ctx context.Context, name string) (*domain.StudentGroup, error) {
	query := `SELECT group_id, name FROM student_groups WHERE name = $1`
	res := m.db.QueryRow(ctx, query, name)
	group := &domain.StudentGroup{}
	err := res.Scan(&group.GroupID, &group.Name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("Unknown group name")
	}
	return group, nil
}

func (m *StudentGroupManager) GetAll(ctx context.Context) ([]domain.StudentGroup, error) {
	query := `SELECT group_id, name FROM student_groups ORDER BY group_id`
	rows, err := m.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []domain.StudentGroup
	for rows.Next() {
		var group domain.StudentGroup
		err := rows.Scan(&group.GroupID, &group.Name)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// --------------------------------------------------------------------------------------------------------------------
