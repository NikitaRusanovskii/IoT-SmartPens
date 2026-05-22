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
