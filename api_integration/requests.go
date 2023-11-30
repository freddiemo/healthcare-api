package api_integration

import "time"

type DiagnosticRequest struct {
	DateTime   *time.Time `binding:"required" json:"dateTime"`
	Diagnostic string     `binding:"required,min=10" json:"diagnostic"`
	PatientID  uint       `binding:"required" json:"patientID"`
}

type DiagnosticRequest2 struct {
	DateTime   string `json:"dateTime,omitempty"`
	Diagnostic string `binding:"required,min=10" json:"diagnostic"`
	PatientID  uint   `binding:"required" json:"patientID"`
}
