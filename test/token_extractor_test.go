package test

import (
	"net/http"
	"testing"

	"github.com/petersonsalme/golang-rest-api/middleware"
)

// TestGivenValidAuthorizationMustReturnToken extract token
func TestGivenValidAuthorizationMustReturnToken(t *testing.T) {
	req := new(http.Request)
	req.Header = map[string][]string{
		"Authorization": {"Bearer 123456789ABC"},
	}

	token := middleware.ExtractToken(req)

	if token != "123456789ABC" {
		t.Fail()
	}
}

// TestGivenInvalidAuthorizationMustReturnEmpty extract token
func TestGivenInvalidAuthorizationMustReturnEmpty(t *testing.T) {
	req := new(http.Request)
	req.Header = map[string][]string{}

	token := middleware.ExtractToken(req)

	if token != "" {
		t.Fail()
	}
}
