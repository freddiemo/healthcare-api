package api_integration

import "time"

type Response struct {
	Message string `json:"message"`
}

type DiagnosticResponse struct {
	ID         uint
	DateTime   *time.Time
	Diagnostic string
	PatientID  uint
}

type PatientResponse struct {
	ID         uint
	FirstName  string
	LastName   string
	Email      string
	DocumentID string
	Phone      string
}

type DiagnosticResponseWithPatient struct {
	ID         uint
	DateTime   *time.Time
	Diagnostic string
	Patient    PatientResponse
}
