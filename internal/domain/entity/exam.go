package entity

import "time"

type Exam struct {
	ID           string
	ExamTypeID   string
	DeviceID     string
	DepartmentID string
	OperatorID   *string
	PhysicianID  *string
	ProbeID      *string
	StartedAt    time.Time
	Status       string
}