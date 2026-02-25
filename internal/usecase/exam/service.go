package exam

import (
	"context"
	"time"

	"github.com/google/uuid"
	"sert_exam_backend/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, exam *entity.Exam) error
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateFromESP(ctx context.Context, deviceID string, departmentID string) error {
	exam := &entity.Exam{
		ID:           uuid.New().String(),
		DeviceID:     deviceID,
		DepartmentID: departmentID,
		StartedAt:    time.Now(),
		Status:       "COMPLETED",
	}

	return s.repo.Create(ctx, exam)
}