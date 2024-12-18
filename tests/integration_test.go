package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"romanum/internal/handlers"
	"testing"
)

func TestConversionHandler(t *testing.T) {
	tests := []struct {
		name       string
		rangeStart string
		rangeEnd   string
		system     string
		wantStatus int
		wantBody   string
	}{
		{"Valid Range Roman", "12", "15", "roman", http.StatusOK, `[{"number":12,"converted":"XII"},{"number":13,"converted":"XIII"},{"number":14,"converted":"XIV"},{"number":15,"converted":"XV"}]`},
		{"Valid Range Binary", "10", "12", "binary", http.StatusOK, `[{"number":10,"converted":"1010"},{"number":11,"converted":"1011"},{"number":12,"converted":"1100"}]`},
		{"Valid Range Hexadecimal", "15", "17", "hexadecimal", http.StatusOK, `[{"number":15,"converted":"F"},{"number":16,"converted":"10"},{"number":17,"converted":"11"}]`},
		{"Missing Range", "blah", "", "", http.StatusBadRequest, "Params 'start' and 'end' needed\n"},
		{"Invalid Range", "15", "10", "", http.StatusBadRequest, "Invalid range. Ensure 1 <= start <= end <= 3999\n"},
		{"Out of Range Start", "0", "10", "", http.StatusBadRequest, "Invalid range. Ensure 1 <= start <= end <= 3999\n"},
		{"Out of Range End", "10", "4000", "", http.StatusBadRequest, "Invalid range. Ensure 1 <= start <= end <= 3999\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/convert?start="+tc.rangeStart+"&end="+tc.rangeEnd+"&system="+tc.system, nil)
			res := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.ConversionHandler)

			handler.ServeHTTP(res, req)

			if status := res.Code; status != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.wantStatus)
			}

			if tc.wantStatus == http.StatusOK {
				var gotBody []handlers.ConversionResult
				json.Unmarshal(res.Body.Bytes(), &gotBody)
				var wantBody []handlers.ConversionResult
				json.Unmarshal([]byte(tc.wantBody), &wantBody)

				if len(gotBody) != len(wantBody) {
					t.Errorf("handler returned unexpected body: got %v want %v", gotBody, wantBody)
				}
				for i := range gotBody {
					if gotBody[i] != wantBody[i] {
						t.Errorf("handler returned unexpected body: got %v want %v", gotBody, wantBody)
					}
				}
			} else {
				if res.Body.String() != tc.wantBody {
					t.Errorf("handler returned unexpected body: got %v want %v", res.Body.String(), tc.wantBody)
				}
			}
		})
	}
}
