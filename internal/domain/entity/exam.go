package entity

import "time"

type Exam struct {
	ID           string
	DeviceID     string
	DepartmentID string
	StartedAt    time.Time
	FinishedAt   *time.Time
	Status       string
}