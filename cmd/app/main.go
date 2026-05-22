package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"smartPens/internal/repository"
	"smartPens/internal/sender"
	"smartPens/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in environment")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	lessonRepo := repository.NewLessonRepository(pool)
	workRepo := repository.NewWorkRepository(pool)

	learningService := service.NewLearningService(workRepo, lessonRepo)

	teacherID := uuid.New()
	groupID := 1
	subjectID := 1
	room := 101
	date := time.Now().UTC()

	lessonID, err := learningService.StartLesson(ctx, teacherID, groupID, subjectID, room, date)
	if err != nil {
		log.Fatalf("Failed to start lesson: %v", err)
	}
	log.Printf("Lesson started with ID: %d", lessonID)

	studentCount := 3
	err = learningService.CompleteLesson(ctx, lessonID, studentCount)
	if err != nil {
		log.Fatalf("Failed to complete lesson: %v", err)
	}
	log.Printf("Lesson %d completed with %d works", lessonID, studentCount)

	senderInstance := &sender.Sender{
		URL:    "http://localhost:8080/api/lesson-with-works",
		DB:     lessonRepo,
		Client: http.DefaultClient,
	}

	if err := senderInstance.SendWorksTo(ctx, lessonID); err != nil {
		log.Fatalf("Failed to send works: %v", err)
	}
	log.Println("Works successfully sent to remote endpoint")
}
