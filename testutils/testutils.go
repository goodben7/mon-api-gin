package testutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertJSONResponse(t *testing.T, w *httptest.ResponseRecorder, expectedCode int, expectedBody interface{}) {
	t.Helper()
	assert.Equal(t, expectedCode, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Failed to parse JSON response")
	}

	expectedJSON, _ := json.Marshal(expectedBody)
	var expectedMap map[string]interface{}
	json.Unmarshal(expectedJSON, &expectedMap)

	assert.Equal(t, expectedMap, response)
}

func NewJSONRequest(method, url string, body interface{}) *http.Request {
	jsonBody, _ := json.Marshal(body)
	return httptest.NewRequest(method, url, strings.NewReader(string(jsonBody)))
}
