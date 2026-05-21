package service

import (
	"math/rand/v2"
	"smartPens/internal/domain"
)

func GenerateDataFromStudentsPen(lessonID int,
	studentCnt int) ([]domain.Work, error) {
	works := []domain.Work{}
	for studentID := range studentCnt {
		var points domain.PointList
		startPosition := [3]int8{10, 20, 30}
		points = append(points, startPosition)
		for range 999 {
			currentPosition := [3]int8{int8(rand.IntN(3) - 1),
				int8(rand.IntN(3) - 1),
				int8(rand.IntN(3) - 1)}
			points = append(points, currentPosition)
		}
		works = append(works, domain.Work{LessonID: lessonID,
			StudentID: studentID,
			Data:      points})
	}
	return works, nil
}
