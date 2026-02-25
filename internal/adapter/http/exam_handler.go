package http

import (
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
	r.Post("/esp/exam", h.CreateExam)
}

func (h *ExamHandler) CreateExam(c *fiber.Ctx) error {
	type Request struct {
		DeviceID     string `json:"deviceId"`
		DepartmentID string `json:"departmentId"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	err := h.service.CreateFromESP(c.Context(), req.DeviceID, req.DepartmentID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot create exam"})
	}

	return c.JSON(fiber.Map{"message": "exam created"})
}