package kairos

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// TestingInterface - General Interface for Testing.
type TestingInterface interface {
	Description(string, Before, Test, After, *testing.T)
}

// Before is all excetion before execute test
type Before struct {
	// Method. e.g. GET, POST, PUT, DELETE
	Method string
	// URL e.g: /api/book/:id
	URL string
	// Body is the json that you will be send.
	// e.g. { "name": "john" }
	Body map[string]interface{}
	// BeforeFunc configuration before run test, such as external agents instance, set Headeres, etc.
	// e.g.: request.Header.Set("Content-Type", "application/json")
	// Default: aplication/json was set.
	BeforeFunc func(*http.Request)
}

// Test is a struct about Testing execution
type Test struct {
	// URL e.g: /api/video/12abb553
	URL string
	// Controller that you would like to test
	Controller func(*fiber.Ctx) error
}

// After Testing is the response of testing
type After struct {
	// Function that is exec when finish testing, recive the response.
	AfterFunc func(*http.Response)
	// StatusCode that you would like to recibe.
	StatusCode int
	// JSON that you would like to recibe for testing.
	Want map[string]interface{}
}
