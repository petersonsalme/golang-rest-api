package test

import (
	"testing"

	"github.com/petersonsalme/golang-rest-api/middleware"
)

// Test test
func TestShouldCreateToken(t *testing.T) {
	_, err := middleware.CreateToken(123)

	if err != nil {
		t.Fail()
	}
}
