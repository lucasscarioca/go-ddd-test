package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lucasscarioca/music-stash-server/configs"
	"github.com/lucasscarioca/music-stash-server/routes"
	"github.com/stretchr/testify/require"
)

// executeRequest, creates a new ResponseRecorder
// then executes the resquest by calling serveHTTP in the router
// after which the handler writes the response to the response recorder
// which we can inspect.
func executeRequest(req *http.Request, s chi.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestHelloWorld(t *testing.T) {
	err := configs.Load()
	if err != nil {
		t.Errorf("Could not initialize environment variables: %s", err.Error())
	}
	s := chi.NewRouter()

	s.Route("/api", routes.Mount)

	// Create a New Request
	req, _ := http.NewRequest("GET", "/api/users", nil)

	// Execute Request
	response := executeRequest(req, s)

	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)

	// We can use testify/require to assert values, as it is more convenient
	require.Equal(t, "Hello World!", response.Body.String())
}
