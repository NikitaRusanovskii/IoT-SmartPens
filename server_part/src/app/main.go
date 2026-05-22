package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"IoT-SmartPens/server_part/internal/domain"
	httpHandler "IoT-SmartPens/server_part/internal/http"
	"IoT-SmartPens/server_part/internal/repository"
	"IoT-SmartPens/server_part/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	lessonService  *service.LessonService
	workService    *service.WorkService
	studentService *service.StudentService
	teacherService *service.TeacherService
	groupService   *service.StudentGroupService
	subjectService *service.SubjectService
)

func saveLessonWithWorks(lesson *domain.Lesson, works []domain.Work) error {
	ctx := context.Background()
	exists, err := teacherService.ExistsByID(ctx, lesson.TeacherID)
	if err != nil {
		return err
	}
	if !exists {
		defaultTeacher := domain.Teacher{
			TeacherID: lesson.TeacherID,
			FIO: domain.Fio{
				FirstName:  "Unknown",
				MiddleName: "",
				LastName:   "Teacher",
			},
			SubjectID: lesson.SubjectID,
		}
		if err := teacherService.Register(ctx, defaultTeacher); err != nil {
			log.Printf("Failed to create default teacher: %v", err)
			return err
		}
		log.Printf("Created default teacher with ID: %s", lesson.TeacherID)
	}
	exists, err = groupService.ExistsByID(ctx, lesson.GroupID)
	if err != nil {
		return err
	}
	if !exists {
		defaultGroup := domain.StudentGroup{
			GroupID: lesson.GroupID,
			Name:    "Default Group",
		}
		if err := groupService.Register(ctx, defaultGroup); err != nil {
			log.Printf("Failed to create default group: %v", err)
			return err
		}
		log.Printf("Created default group with ID: %d", lesson.GroupID)
	}

	exists, err = subjectService.ExistsByID(ctx, lesson.SubjectID)
	if err != nil {
		return err
	}
	if !exists {
		defaultSubject := domain.Subject{
			SubjectID: lesson.SubjectID,
			Name:      "Default Subject",
		}
		if err := subjectService.Register(ctx, defaultSubject); err != nil {
			log.Printf("Failed to create default subject: %v", err)
			return err
		}
		log.Printf("Created default subject with ID: %d", lesson.SubjectID)
	}
	if err := lessonService.Register(ctx, *lesson); err != nil && err.Error() != "already registered" {
		return err
	}
	for _, work := range works {
		exists, err := studentService.ExistsByID(ctx, work.StudentID)
		if err != nil {
			return err
		}
		if !exists {
			defaultStudent := domain.Student{
				StudentID: work.StudentID,
				FIO: domain.Fio{
					FirstName:  "Unknown",
					MiddleName: "",
					LastName:   "Student",
				},
				GroupID: lesson.GroupID,
			}
			if err := studentService.Register(ctx, defaultStudent); err != nil {
				log.Printf("Failed to create default student: %v", err)
				return err
			}
			log.Printf("Created default student with ID: %s", work.StudentID)
		}
		if err := workService.Register(ctx, work); err != nil {
			return err
		}
	}
	log.Printf("Saved lesson %d with %d works", lesson.LessonID, len(works))
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbURL := os.Getenv("DATABASE_URL")

	ctx := context.Background()
	connManager := repository.NewConnectionManager(nil)

	dbPool, err := connManager.Connect(ctx, dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer connManager.Disconnect()

	log.Println("Database connected successfully")

	lessonRepo, _ := repository.NewLessonManager(dbPool)
	workRepo, _ := repository.NewWorkManager(dbPool)
	studentRepo, _ := repository.NewStudentManager(dbPool)
	teacherRepo, _ := repository.NewTeacherManager(dbPool)
	groupRepo, _ := repository.NewStudentGroupManager(dbPool)
	subjectRepo, _ := repository.NewSubjectManager(dbPool)

	lessonService = service.NewLessonService(lessonRepo)
	workService = service.NewWorkService(workRepo)
	studentService = service.NewStudentService(studentRepo)
	teacherService = service.NewTeacherService(teacherRepo)
	groupService = service.NewStudentGroupService(groupRepo)
	subjectService = service.NewSubjectService(subjectRepo)

	r := gin.Default()

	r.POST("/api/lesson-with-works", httpHandler.CreateLessonWithWorksHandler(saveLessonWithWorks))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	port := "8080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
