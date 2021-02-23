package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/KummerCompany/kairos"
	"github.com/KummerCompany/kairos/pkg"
	"github.com/gofiber/fiber/v2"
)

type fiberTesting struct{}

// NewFiberTestAPI is a fiber constructor for TESTING
func NewFiberTestAPI() kairos.TestingInterface {
	return &fiberTesting{}
}

func (*fiberTesting) Description(escenario string, before kairos.Before, test kairos.Test, after kairos.After, t *testing.T) {

	c := pkg.ColorsPrint{}

	c.Cyan("\n" + escenario)

	app := fiber.New()

	switch before.Method {
	case "GET":
		app.Get(before.URL, test.Controller)
	case "POST":
		app.Post(before.URL, test.Controller)
	case "PUT":
		app.Put(before.URL, test.Controller)
	case "DELETE":
		app.Delete(before.URL, test.Controller)
	}

	body, _ := json.Marshal(before.Body)

	request := httptest.NewRequest(before.Method, before.URL, bytes.NewBuffer(body))

	request.Header.Set("Content-Type", "application/json")

	if before.BeforeFunc != nil {
		before.BeforeFunc(request)
	}

	resp, err := app.Test(request)

	if err != nil {
		t.Fail()
	}

	if after.AfterFunc != nil {
		after.AfterFunc(resp)
	}

	var got map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&got)

	gotB, _ := json.Marshal(got)
	wantB, _ := json.Marshal(after.Want)

	gotS := string(gotB)
	wantS := string(wantB)

	if gotS != wantS {
		c.Red("\n\tI recibe " + gotS)
		c.Red("\tbut I want " + wantS)
		t.Fail()
	} else {
		c.Green("\n\tI recibe " + gotS)
	}

	if resp.StatusCode != after.StatusCode {
		c.Red("\n\tI recibe " + fmt.Sprint(resp.StatusCode))
		c.Red("\tbut I want " + fmt.Sprint(after.StatusCode))
		t.Fail()
	} else {
		c.Green("\n\tI recibe " + fmt.Sprint(resp.StatusCode))
	}
}
