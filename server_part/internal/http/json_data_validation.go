package http

import (
	"IoT-SmartPens/server_part/internal/domain"
	"errors"

	"github.com/google/uuid"
)

// Works Validation ---------------------------------------------------------------------------------------------------

func ValidateWorksData(works []domain.Work) error {
	for i, work := range works {
		if work.StudentID == uuid.Nil {
			return errors.New("student_id is required at index " + string(rune(i)))
		}
		if work.Data == nil {
			return errors.New("coords cannot be nil at index " + string(rune(i)))
		}
		for j, coord := range work.Data {
			if coord.Xcoord < 0 || coord.Xcoord > 10000 {
				return errors.New("invalid Xcoord at work " + string(rune(i)) + ", coord " + string(rune(j)))
			}
			if coord.Ycoord < 0 || coord.Ycoord > 10000 {
				return errors.New("invalid Ycoord at work " + string(rune(i)) + ", coord " + string(rune(j)))
			}
			if coord.Pressure < 0 || coord.Pressure > 1024 {
				return errors.New("invalid Pressure at work " + string(rune(i)) + ", coord " + string(rune(j)))
			}
		}
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// Teachers and Students Validation -----------------------------------------------------------------------------------

func ValidateTeachersStudentsData(teachers []domain.Teacher, students []domain.Student) error {
	teacherIDs := make(map[uuid.UUID]bool)
	for i, teacher := range teachers {
		if teacherIDs[teacher.TeacherID] {
			return errors.New("duplicate teacher_id found at index " + string(rune(i)))
		}
		teacherIDs[teacher.TeacherID] = true

		if teacher.FIO.FirstName == "" {
			return errors.New("teacher first_name is required at index " + string(rune(i)))
		}
		if teacher.FIO.LastName == "" {
			return errors.New("teacher last_name is required at index " + string(rune(i)))
		}
		if teacher.SubjectID <= 0 {
			return errors.New("teacher subject_id must be positive at index " + string(rune(i)))
		}
	}
	studentIDs := make(map[uuid.UUID]bool)
	for i, student := range students {
		if studentIDs[student.StudentID] {
			return errors.New("duplicate student_id found at index " + string(rune(i)))
		}
		studentIDs[student.StudentID] = true

		if student.FIO.FirstName == "" {
			return errors.New("student first_name is required at index " + string(rune(i)))
		}
		if student.FIO.LastName == "" {
			return errors.New("student last_name is required at index " + string(rune(i)))
		}
		if student.GroupID <= 0 {
			return errors.New("student group_id must be positive at index " + string(rune(i)))
		}
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------
