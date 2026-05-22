package http

import (
	"net/http"

	"IoT-SmartPens/server_part/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// JSON -> structs full logic -----------------------------------------------------------------------------------------

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

func CreateTeachersStudentsHandler(
	saveFunc func([]domain.Teacher, []domain.Student) error,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		teachers, students, err := ParseTeachersStudentsData(c)
		if err != nil {
			sendError(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := ValidateTeachersStudentsData(teachers, students); err != nil {
			sendError(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := saveFunc(teachers, students); err != nil {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":         "success",
			"teachers_saved": len(teachers),
			"students_saved": len(students),
			"message":        "Teachers and students saved successfully",
		})
	}
}
