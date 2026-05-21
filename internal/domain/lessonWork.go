package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

/*type StudentWork struct {
	StudentID uuid.UUID        `json:"student_id"`
	Points    domain.PointList `json:"coords"`
}*/

type WorksRequest struct {
	LessonID  int               `json:"lesson_id"`
	TeacherID uuid.UUID         `json:"teacher_id"`
	Date      time.Time         `json:"date"`
	GroupID   int               `json:"group_id"`
	SubjectID int               `json:"subject_id"`
	Room      int               `json:"room"`
	Works     []json.RawMessage `json:"works"`
}
