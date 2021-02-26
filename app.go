package kairos

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/KummerCompany/kairos/pkg"
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
	// URL that you will testing. e.g. /api/videos
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

// Assert struct
type Assert struct{}

// Equal : Compare two values, that you would like to recibe and you recibe.
func (*Assert) Equal(got, want interface{}, t *testing.T) {
	c := pkg.ColorsPrint{}

	typeGot := reflect.TypeOf(got)
	typeWant := reflect.TypeOf(want)
	valueGot := reflect.ValueOf(got)
	valueWant := reflect.ValueOf(want)

	if typeGot != typeWant {
		c.Red("\n\tI recibe")
		fmt.Println(typeGot)
		c.Red("\n\tBut I want")
		fmt.Println(typeWant)
		t.Fail()
		return
	}
	if valueGot != valueWant {
		c.Red("\n\tI recibe")
		fmt.Println(valueGot)
		c.Red("\n\tBut I want")
		fmt.Println(valueWant)
		t.Fail()
		return
	}

	c.Green("\n\tBoth values are equals")
}
