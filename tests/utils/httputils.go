package utilstests

import (
	"net/http"
	"net/http/httptest"

	"gorm.io/gorm"

	"github.com/freddiemo/healthcare-api/api"
)

func ExecuteRequest(req *http.Request, db *gorm.DB) *httptest.ResponseRecorder {
	r := api.GetServer(db)
	nr := httptest.NewRecorder()

	r.ServeHTTP(nr, req)

	return nr
}
