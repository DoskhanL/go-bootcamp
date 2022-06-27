package weather

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Weather
	expectedError error
}

func TestGetWeather(t *testing.T) {

	tests := []Tests{
		{
			name: "basic-request",
			server: httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, _ *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"city": "Almaty, Kazakhstan", "forecast":"sunny"}`))
				},
			)),
			response:      &Weather{City: "Almaty, Kazakhstan", Forecast: "sunny"},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer test.server.Close()
			resp, err := GetWeather(test.server.URL)

			if !reflect.DeepEqual(test.response, resp) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}

			if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error FAILED: expected %v, got %v", test.expectedError, err)
			}
		})
	}
}
