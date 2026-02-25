package exam

import (
	"context"

	"github.com/google/uuid"
	"sert_exam_backend/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, e *entity.Exam) error
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, req CreateExamRequest) error {

	exam := &entity.Exam{
		ID:           uuid.New().String(),
		ExamTypeID:   req.ExamTypeID,
		DeviceID:     req.DeviceID,
		DepartmentID: req.DepartmentID,
		OperatorID:   req.OperatorID,
		PhysicianID:  req.PhysicianID,
		ProbeID:      req.ProbeID,
		StartedAt:    req.StartedAt,
		Status:       "COMPLETED",
	}

	return s.repo.Create(ctx, exam)
}