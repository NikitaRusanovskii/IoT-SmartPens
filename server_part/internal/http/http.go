package http

import (
	"errors"
	"net/http"
	"time"

	"IoT-SmartPens/server_part/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
		LessonID:  rawData.LessonID,
		TeacherID: rawData.TeacherID,
		Date:      rawData.Date,
		GroupID:   rawData.GroupID,
		SubjectID: rawData.SubjectID,
		Room:      rawData.Room,
	}
	var works []domain.Work
	sum := 0
	for i, wd := range rawData.Works {
		sum += i
		work := domain.Work{
			WorkID:    0, // Будет сгенерирован при вставке в БД
			LessonID:  rawData.LessonID,
			StudentID: wd.StudentID,
			Data:      wd.Coords,
			Mark:      0, // По умолчанию не оценено
		}
		works = append(works, work)
	}

	return lesson, works, nil
}

func ValidateWorksData(works []domain.Work) error {
	if len(works) == 0 {
		return errors.New("works cannot be empty")
	}
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

func sendError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func CreateLessonWithWorksHandler(saveFunc func(*domain.Lesson, []domain.Work) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		lesson, works, err := ParseLessonWithWorksData(c)
		if err != nil {
			sendError(c, http.StatusBadRequest, err.Error())
			return
		}
		if lesson.LessonID == 0 {
			sendError(c, http.StatusBadRequest, "lesson_id is required")
			return
		}
		if lesson.TeacherID == uuid.Nil {
			sendError(c, http.StatusBadRequest, "teacher_id is required")
			return
		}
		if err := ValidateWorksData(works); err != nil {
			sendError(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := saveFunc(lesson, works); err != nil {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":      "success",
			"lesson_id":   lesson.LessonID,
			"works_count": len(works),
		})
	}
}
