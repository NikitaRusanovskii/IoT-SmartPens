package http

import (
	"IoT-SmartPens/server_part/internal/domain"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// --------------------------------------------------------------------------------------------------------------------

// ParseLessonWithWorksData парсит JSON формата:
//
//	{
//	    "lesson_id": 1,
//	    "teacher_id": "123e4567-e89b-12d3-a456-426614174000",
//	    "date": "2024-01-15T10:00:00Z",
//	    "group_id": 10,
//	    "subject_id": 5,
//	    "room": 301,
//	    "works_data": [
//	        {"student_id": "223e4567-e89b-12d3-a456-426614174001", "coords": [{"Xcoord": 10, "Ycoord": 20, "Pressure": 512}]},
//	        {"student_id": "323e4567-e89b-12d3-a456-426614174002", "coords": [{"Xcoord": 15, "Ycoord": 25, "Pressure": 520}]},
//	        {"student_id": "423e4567-e89b-12d3-a456-426614174003", "coords": []}
//	    ]
//	}
func ParseLessonWithWorksData(c *gin.Context) (*domain.Lesson, []domain.Work, error) {
	var rawData struct {
		LessonID  int       `json:"lesson_id"`
		TeacherID uuid.UUID `json:"teacher_id"`
		Date      time.Time `json:"date"`
		GroupID   int       `json:"group_id"`
		SubjectID int       `json:"subject_id"`
		Room      int       `json:"room"`
		Works     []struct {
			StudentID uuid.UUID       `json:"student_id"`
			Coords    []domain.Coords `json:"coords"`
		} `json:"works"`
	}
	if err := c.ShouldBindJSON(&rawData); err != nil {
		return nil, nil, errors.New("failed to parse JSON: " + err.Error())
	}
	lesson := &domain.Lesson{
		LessonID:  0, // Будет сгенерирован при вставке в БД
		TeacherID: rawData.TeacherID,
		Date:      rawData.Date,
		GroupID:   rawData.GroupID,
		SubjectID: rawData.SubjectID,
		Room:      rawData.Room,
	}
	var works []domain.Work
	sum := 0
	for i, wd := range rawData.Works {
		work := domain.Work{
			WorkID:    0, // Будет сгенерирован при вставке в БД
			LessonID:  rawData.LessonID,
			StudentID: wd.StudentID,
			Data:      wd.Coords,
			Mark:      0, // По умолчанию не оценено
		}
		works = append(works, work)
		sum += i
	}

	return lesson, works, nil
}

// --------------------------------------------------------------------------------------------------------------------
