package internal

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// Teacher Service ----------------------------------------------------------------------------------------------------

type TeacherService struct {
	m *TeacherManager
}

func (ts *TeacherService) Register(ctx context.Context, t Teacher) error {
	isExists, err := ts.m.ExistsByID(ctx, t.TeacherID)
	if err != nil {
		return err
	}
	if isExists {
		return errors.New("Teacher already registered")
	}
	err = ts.m.Insert(ctx, t)
	return err
}
func (ts *TeacherService) UpdateSubject(ctx context.Context, id uuid.UUID, subjID int) error {
	isExists, err := ts.m.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !isExists {
		return errors.New("Teacher does not exists")
	}
	teac, err := ts.m.GetByID(ctx, id)
	if err != nil {
		return err
	}
	teac.SubjectID = subjID
	err = ts.m.Insert(ctx, *teac)
	return err
}
func (ts *TeacherService) Remove(ctx context.Context, id uuid.UUID) error {
	err := ts.m.DeleteByID(ctx, id)
	return err
}

// --------------------------------------------------------------------------------------------------------------------

// Student Service ----------------------------------------------------------------------------------------------------

type StudentService struct {
	m *StudentManager
}

func (ts *StudentService) Register(ctx context.Context, t Student) error {
	isExists, err := ts.m.ExistsByID(ctx, t.StudentID)
	if err != nil {
		return err
	}
	if isExists {
		return errors.New("Student already registered")
	}
	err = ts.m.Insert(ctx, t)
	return err
}
func (ts *StudentService) UpdateGroup(ctx context.Context, id uuid.UUID, groupID int) error {
	isExists, err := ts.m.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !isExists {
		return errors.New("Student does not exists")
	}
	stud, err := ts.m.GetByID(ctx, id)
	if err != nil {
		return err
	}
	stud.GroupID = groupID
	err = ts.m.Insert(ctx, *stud)
	return err
}
func (ts *StudentService) Remove(ctx context.Context, id uuid.UUID) error {
	err := ts.m.DeleteByID(ctx, id)
	return err
}

// --------------------------------------------------------------------------------------------------------------------
