package service

import (
	"math/rand/v2"
	"smartPens/internal/domain"

	"github.com/google/uuid"
)

func GenerateDataFromStudentsPen(lessonID int,
	studentCnt int) ([]domain.Work, error) {
	works := []domain.Work{}
	for range studentCnt {
		studentID := uuid.New()
		var points []domain.Point
		startPosition := domain.Point{X: 10, Y: 20, Pressure: 30}
		points = append(points, startPosition)
		for range 3 {
			x := int8(rand.IntN(3) - 1)
			y := int8(rand.IntN(3) - 1)
			pressure := int8(rand.IntN(100))
			currentPosition := domain.Point{X: x, Y: y, Pressure: pressure}
			points = append(points, currentPosition)
		}
		works = append(works, domain.Work{LessonID: lessonID,
			StudentID: studentID,
			Data:      points})
	}
	return works, nil
}
