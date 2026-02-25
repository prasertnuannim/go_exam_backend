package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	httpadapter "sert_exam_backend/internal/adapter/http"
	"sert_exam_backend/internal/adapter/persistence"
	"sert_exam_backend/internal/infrastructure/config"
	"sert_exam_backend/internal/infrastructure/database"
	"sert_exam_backend/internal/usecase/exam"
)

func main() {
	cfg := config.Load()
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal("❌ Cannot connect DB:", err)
	}
	if err := db.AutoMigrate(&persistence.ExamModel{}); err != nil {
		log.Fatal("❌ Cannot migrate DB:", err)
	}

	app := fiber.New()
	repo := persistence.NewExamRepository(db)
	service := exam.NewService(repo)
	handler := httpadapter.NewExamHandler(service)
	handler.Register(app)

	log.Printf("🚀 Server running on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
