package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
)

var testDiagnostic model.Diagnostic = model.Diagnostic{
	DateTime:   &DATETIME,
	Diagnostic: DIAGNOSTIC,
	PatientID:  PATIENTID,
}

var _ = Describe("Diagnostics Service", func() {

	BeforeEach(func() {
		diagnosticsService.DeleteAll()
	})

	AfterEach(func() {
		diagnosticsService.DeleteAll()
	})

	Describe("Fetching existing diagnostics", func() {
		Context("If there are no diagnostics in the database", func() {

			It("should return an empty list", func() {
				diagnostics, _ := diagnosticsService.Find(&model.DiagnosticQuery{})

				Ω(diagnostics).Should(BeEmpty())
			})
		})

		Context("If therre is a diagnostic in the database", func() {
			BeforeEach(func() {
				diagnosticsService.Save(&testDiagnostic)
			})

			AfterEach(func() {
				diagnosticList, _ := diagnosticsService.Find(&model.DiagnosticQuery{})
				diagnostic := diagnosticList[0]
				diagnosticsService.Delete(diagnostic.ID)
			})

			It("should return at least one element", func() {
				diagnosticList, _ := diagnosticsService.Find(&model.DiagnosticQuery{})

				Ω(diagnosticList).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				diagnosticList, _ := diagnosticsService.Find(&model.DiagnosticQuery{})
				firstDiagnostic := diagnosticList[0]
				Ω(firstDiagnostic.Diagnostic).Should(Equal(DIAGNOSTIC))
				Ω(firstDiagnostic.PatientID).Should(Equal(PATIENTID))
			})
		})

	})

})
