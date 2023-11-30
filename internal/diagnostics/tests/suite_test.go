package tests

import (
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	dbsetup "github.com/freddiemo/healthcare-api/db"
	"github.com/freddiemo/healthcare-api/internal/diagnostics/repository"
	"github.com/freddiemo/healthcare-api/internal/diagnostics/service"
	utilstests "github.com/freddiemo/healthcare-api/tests/utils"
)

const (
	PATIENTID uint = 1
)

var (
	diagnosticsRepository repository.DiagnosticsRepository
	diagnosticsService    service.DiagnosticsService
	db                    *gorm.DB
	DATETIME              = time.Now()
	DIAGNOSTIC            = faker.Paragraph()
)

func TestDiagnosticService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Diagnostics Service Test Suite")
}

var _ = BeforeSuite(func() {
	params := utilstests.LoadEnv()
	db = dbsetup.Init(params)

	diagnosticsRepository = repository.NewDiagnosticsRepository(db)
	diagnosticsService = service.NewDiagnosticsService(diagnosticsRepository)
})
