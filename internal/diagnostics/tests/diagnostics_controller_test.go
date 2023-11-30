package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/freddiemo/healthcare-api/internal/diagnostics/model"
	utilstests "github.com/freddiemo/healthcare-api/tests/utils"
)

var _ = Describe("DiagnosticController", func() {

	var diagnostic *model.Diagnostic
	var apiVersion = "/v1"

	BeforeEach(func() {
		diagnosticsService.DeleteAll()

		diagnostic = &model.Diagnostic{
			DateTime:   &DATETIME,
			Diagnostic: DIAGNOSTIC,
			PatientID:  PATIENTID,
		}
		diagnosticsService.Save(diagnostic)
	})

	AfterEach(func() {
		diagnosticsService.DeleteAll()
	})

	Describe("List", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/diagnostics/", apiVersion), nil)
			response = utilstests.ExecuteRequest(req, db)
		})

		Context("should list all diagnostics", func() {

			It("status code should be 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				l, _ := utilstests.DeserializeList(response.Body.String())
				Expect(len(l)).To(Equal(1))
				for _, m := range l {
					Expect(m["ID"]).To(BeEquivalentTo(diagnostic.ID))
					Expect(m["Diagnostic"]).To(Equal(DIAGNOSTIC))
				}
			})

		})

	})

	Describe("Create", func() {
		var response *httptest.ResponseRecorder
		var payload []byte

		BeforeEach(func() {
			diagnostic = &model.Diagnostic{
				DateTime:   &DATETIME,
				Diagnostic: DIAGNOSTIC,
				PatientID:  PATIENTID,
			}
			payload, _ = json.Marshal(diagnostic)
		})

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/diagnostics/", apiVersion), bytes.NewBuffer(payload))
			response = utilstests.ExecuteRequest(req, db)
		})

		Context("should create a monster correctly", func() {

			It("status code should be 201", func() {
				Expect(response.Code).To(Equal(201))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				m, _ := utilstests.Deserialize(response.Body.String())
				fmt.Println("m:", m)
				Expect(m["ID"]).ToNot(BeNil())
				Expect(m["diagnostic"]).To(Equal(DIAGNOSTIC))
				Expect(m["patientID"]).To(BeEquivalentTo(PATIENTID))
			})

		})

		Context("should get an error when create a diagnostics without patient", func() {

			BeforeEach(func() {
				diagnostic = &model.Diagnostic{
					DateTime:   &DATETIME,
					Diagnostic: DIAGNOSTIC,
				}
				payload, _ = json.Marshal(diagnostic)
			})

			It("status code should be 400", func() {
				Expect(response.Code).To(Equal(400))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

		})

	})

})
