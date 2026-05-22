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

// ParseTeachersStudentsData парсит JSON формата:
//
//	{
//	    "teachers": [
//	        {
//	            "teacher_id": "123e4567-e89b-12d3-a456-426614174000",
//	            "fio": {
//	                "first_name": "Иван",
//	                "middle_name": "Петрович",
//	                "last_name": "Сидоров"
//	            },
//	            "subject_id": 5
//	        }
//	    ],
//	    "students": [
//	        {
//	            "student_id": "323e4567-e89b-12d3-a456-426614174002",
//	            "fio": {
//	                "first_name": "Алексей",
//	                "middle_name": "Дмитриевич",
//	                "last_name": "Кузнецов"
//	            },
//	            "group_id": 10
//	        }
//	    ]
//	}
func ParseTeachersStudentsData(c *gin.Context) ([]domain.Teacher, []domain.Student, error) {
	var rawData struct {
		Teachers []struct {
			TeacherID uuid.UUID `json:"teacher_id"`
			FIO       struct {
				FirstName  string `json:"first_name"`
				MiddleName string `json:"middle_name"`
				LastName   string `json:"last_name"`
			} `json:"fio"`
			SubjectID int `json:"subject_id"`
		} `json:"teachers"`
		Students []struct {
			StudentID uuid.UUID `json:"student_id"`
			FIO       struct {
				FirstName  string `json:"first_name"`
				MiddleName string `json:"middle_name"`
				LastName   string `json:"last_name"`
			} `json:"fio"`
			GroupID int `json:"group_id"`
		} `json:"students"`
	}

	if err := c.ShouldBindJSON(&rawData); err != nil {
		return nil, nil, errors.New("failed to parse JSON: " + err.Error())
	}

	teachers := make([]domain.Teacher, 0, len(rawData.Teachers))
	for i, t := range rawData.Teachers {
		if t.TeacherID == uuid.Nil {
			return nil, nil, errors.New("teacher_id is required for teacher at index " + string(rune(i)))
		}
		if t.SubjectID == 0 {
			return nil, nil, errors.New("subject_id is required for teacher at index " + string(rune(i)))
		}

		teacher := domain.Teacher{
			TeacherID: t.TeacherID,
			FIO: domain.Fio{
				FirstName:  t.FIO.FirstName,
				MiddleName: t.FIO.MiddleName,
				LastName:   t.FIO.LastName,
			},
			SubjectID: t.SubjectID,
		}
		teachers = append(teachers, teacher)
	}

	students := make([]domain.Student, 0, len(rawData.Students))
	for i, s := range rawData.Students {
		if s.StudentID == uuid.Nil {
			return nil, nil, errors.New("student_id is required for student at index " + string(rune(i)))
		}
		if s.GroupID == 0 {
			return nil, nil, errors.New("group_id is required for student at index " + string(rune(i)))
		}

		student := domain.Student{
			StudentID: s.StudentID,
			FIO: domain.Fio{
				FirstName:  s.FIO.FirstName,
				MiddleName: s.FIO.MiddleName,
				LastName:   s.FIO.LastName,
			},
			GroupID: s.GroupID,
		}
		students = append(students, student)
	}

	return teachers, students, nil
}

// --------------------------------------------------------------------------------------------------------------------
