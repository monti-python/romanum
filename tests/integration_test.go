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
		wantStatus int
		wantBody   string
	}{
		{"Valid Range", "12", "15", http.StatusOK, `[{"number":12,"roman":"XII"},{"number":13,"roman":"XIII"},{"number":14,"roman":"XIV"},{"number":15,"roman":"XV"}]`},
		{"Missing Range", "blah", "", http.StatusBadRequest, "Params 'start' and 'end' needed\n"},
		{"Invalid Range", "15", "10", http.StatusBadRequest, "Invalid range. Ensure 1 <= start <= end <= 3999\n"},
		{"Out of Range Start", "0", "10", http.StatusBadRequest, "Invalid range. Ensure 1 <= start <= end <= 3999\n"},
		{"Out of Range End", "10", "4000", http.StatusBadRequest, "Invalid range. Ensure 1 <= start <= end <= 3999\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/convert?start="+tc.rangeStart+"&end="+tc.rangeEnd, nil)
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
