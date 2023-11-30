package api_integration

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func parseDiagnosticRQ(ctx *gin.Context) (*DiagnosticRequest, error) {
	diagnosticRQ := DiagnosticRequest{}
	diagnosticRQ2 := DiagnosticRequest2{}
	if errRQ2 := ctx.ShouldBind(&diagnosticRQ2); errRQ2 == nil {
		setDiagnostic(&diagnosticRQ2, &diagnosticRQ)
		if len(diagnosticRQ2.DateTime) == 0 {
			now := time.Now()
			diagnosticRQ.DateTime = &now

			return &diagnosticRQ, nil
		}

		if time, err := stringToDate(diagnosticRQ2.DateTime); err == nil {
			diagnosticRQ.DateTime = time

			return &diagnosticRQ, nil
		}

	} else {
		return nil, errRQ2
	}

	return nil, errors.New("Invalid payload")
}

func setDiagnostic(diagnosticRQ2 *DiagnosticRequest2, diagnosticRQ *DiagnosticRequest) {
	diagnosticRQ.Diagnostic = diagnosticRQ2.Diagnostic
	diagnosticRQ.PatientID = diagnosticRQ2.PatientID
}

func stringToDate(stringDateTime string) (dateTime *time.Time, err error) {
	formats := []string{"2006-01-02T15:04:05Z", "2006-01-02 3:04PM", "2/1/2006 15:04:05", "3.04 2 Jan - 2006"}
	for i := 0; i < len(formats); i++ {
		if dateTime, err := time.Parse(formats[i], stringDateTime); err == nil {
			return &dateTime, nil
		}
	}

	return nil, err
}
