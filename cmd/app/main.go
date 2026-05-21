package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"smartPens/internal/repository"
	"smartPens/internal/service"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	/*

		~/go/bin/migrate -path migrations -database "postgres://iot:iot@localhost:5433/iot_db?sslmode=disable" up

	*/

	if err := godotenv.Load(); err != nil {
		panic("godotenv.Load() error")
	}
	dbURL := os.Getenv("DATABASE_URL")

	ctx := context.Background()

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		panic(err)
	}
	workRepo := repository.NewWorkRepository(db)
	lessonRepo := repository.NewLessonRepository(db)

	// Создание экземпляра сервиса
	learningService := service.NewLearningService(workRepo, lessonRepo)
	// 1. Тестирование StartLesson
	lessonID := 1
	teacherID := 1
	groupID := 10
	subjectID := 5
	room := 301
	date := time.Now()

	err = learningService.StartLesson(ctx, lessonID, teacherID, groupID, subjectID, room, date)
	if err != nil {
		log.Fatalf("StartLesson failed: %v", err)
	}
	fmt.Println("✓ StartLesson executed successfully")

	// 2. Тестирование CompleteLesson (генерирует работы для 5 учеников)
	studentCnt := 5
	err = learningService.CompleteLesson(ctx, lessonID, studentCnt)
	if err != nil {
		log.Fatalf("CompleteLesson failed: %v", err)
	}
	fmt.Println("✓ CompleteLesson executed successfully")

	// 3. Тестирование Delete
	studentID := 3
	err = learningService.Delete(ctx, lessonID, studentID)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	fmt.Println("✓ Delete executed successfully")
}
