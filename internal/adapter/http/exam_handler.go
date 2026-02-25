package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"sert_exam_backend/internal/usecase/exam"
)

type ExamHandler struct {
	service *exam.Service
}

func NewExamHandler(s *exam.Service) *ExamHandler {
	return &ExamHandler{service: s}
}

func (h *ExamHandler) Register(r fiber.Router) {
	r.Post("/api/exams", h.Create)
}

func (h *ExamHandler) Create(c *fiber.Ctx) error {
	var req exam.CreateExamRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	log.Println("📥 New Exam Received")
	log.Printf("DeviceID: %s\n", req.DeviceID)
	log.Printf("ExamTypeID: %s\n", req.ExamTypeID)
	log.Printf("DepartmentID: %s\n", req.DepartmentID)
	log.Printf("StartedAt: %v\n", req.StartedAt)

	if err := h.service.Create(c.Context(), req); err != nil {
		log.Println("❌ Insert failed:", err)
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	log.Println("✅ Exam saved successfully")
	log.Println("------------------------------------------------")

	return c.JSON(fiber.Map{"message": "exam created"})
}
