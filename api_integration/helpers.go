package api_integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getDiagnostics(params map[string]string, path string) ([]DiagnosticResponseWithPatient, error) {
	diagnosticRespList := []DiagnosticResponseWithPatient{}
	requestURL := fmt.Sprintf("%s://%s:%s:%s", params["APP_PROTOCOL"], params["APP_HOST"], params["APP_PORT"], path)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return diagnosticRespList, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return diagnosticRespList, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return diagnosticRespList, err
	}

	if err := json.Unmarshal(resBody, &diagnosticRespList); err != nil {
		return diagnosticRespList, err
	}

	return diagnosticRespList, nil
}

func postDiagnostic(params map[string]string, path string, payload *DiagnosticRequest) (*DiagnosticResponse, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	requestURL := fmt.Sprintf("%s://%s:%s:%s", params["APP_PROTOCOL"], params["APP_HOST"], params["APP_PORT"], path)
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	diagnosticResp := DiagnosticResponse{}
	if err := json.Unmarshal(body, &diagnosticResp); err != nil {
		return nil, err
	}

	return &diagnosticResp, nil
}
