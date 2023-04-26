package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stinekamau/simplebank/token"
)

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name           string
		setupAuth      func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checkResponses func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
