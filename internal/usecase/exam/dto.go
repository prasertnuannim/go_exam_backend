package exam

import "time"

type CreateExamRequest struct {
	ExamTypeID   string     `json:"examTypeId"`
	DeviceID     string     `json:"deviceId"`
	DepartmentID string     `json:"departmentId"`
	OperatorID   *string    `json:"operatorId"`
	PhysicianID  *string    `json:"physicianId"`
	ProbeID      *string    `json:"probeId"`
	StartedAt    time.Time  `json:"startedAt"`
}