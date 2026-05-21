package repository

import (
	"IoT-SmartPens/server_part/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Student & Teacher Managers General Functions -----------------------------------------------------------------------

type tableConfig struct {
	table       string
	id_name     string
	oth_id_name string
}

// TO DO: protection from SQL injection (use pgx.Identifier{...}.Sanitize())

func insertFunctional(ctx context.Context, db *pgxpool.Pool,
	id uuid.UUID, fio domain.Fio, other_id int, tc tableConfig) error {
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
