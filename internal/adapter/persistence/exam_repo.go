package persistence

import (
	"context"
	"time"

	"gorm.io/gorm"
	"sert_exam_backend/internal/domain/entity"
)

type ExamModel struct {
	ID           string    `gorm:"column:id;type:uuid;primaryKey"`
	ExamTypeID   string    `gorm:"column:examTypeId;type:uuid;not null"`
	DeviceID     string    `gorm:"column:deviceId;type:uuid;not null"`
	DepartmentID string    `gorm:"column:departmentId;type:uuid;not null"`
	OperatorID   *string   `gorm:"column:operatorId;type:uuid"`
	PhysicianID  *string   `gorm:"column:physicianId;type:uuid"`
	ProbeID      *string   `gorm:"column:probeId;type:uuid"`
	StartedAt    time.Time `gorm:"column:startedAt;type:timestamp(3);not null"`
	Status       string    `gorm:"column:status;type:ExamStatus;not null"`
}

func (ExamModel) TableName() string {
	return "Exam"
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
		ExamTypeID:   exam.ExamTypeID,
		DeviceID:     exam.DeviceID,
		DepartmentID: exam.DepartmentID,
		OperatorID:   exam.OperatorID,
		PhysicianID:  exam.PhysicianID,
		ProbeID:      exam.ProbeID,
		StartedAt:    exam.StartedAt,
		Status:       exam.Status,
	}
	return r.db.WithContext(ctx).Create(&model).Error
}
