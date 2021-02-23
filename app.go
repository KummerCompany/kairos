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
	// URL that you will testing
	URL string
	// Body is the json that you will be send
	Body map[string]interface{}
	// BeforeFunc configuration before run test. e.g external agents instance, set Headeres, etc
	BeforeFunc func(*http.Request)
}

// Test is a struct about Testing execution
type Test struct {
	// Controller that you would like to test
	Controller func(*fiber.Ctx) error
}

// After Testing is the response of testing
type After struct {
	AfterFunc  func(*http.Response)
	StatusCode int
	Want       map[string]interface{}
}
