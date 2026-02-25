package persistence

import (
	"context"

	"gorm.io/gorm"
	"sert_exam_backend/internal/domain/entity"
)

type ExamModel struct {
	ID           string `gorm:"type:uuid;primaryKey"`
	DeviceID     string
	DepartmentID string
	StartedAt    int64
	Status       string
}

type ExamRepository struct {
	db *gorm.DB
}

func NewExamRepository(db *gorm.DB) *ExamRepository {
	return &ExamRepository{db: db}
}

func (r *ExamRepository) Create(ctx context.Context, exam *entity.Exam) error {
	model := ExamModel{
		ID:           exam.ID,
		DeviceID:     exam.DeviceID,
		DepartmentID: exam.DepartmentID,
		StartedAt:    exam.StartedAt.Unix(),
		Status:       exam.Status,
	}
	return r.db.WithContext(ctx).Create(&model).Error
}