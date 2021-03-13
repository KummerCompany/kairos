package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/KummerCompany/kairos/pkg"
)

// Equal : Compare two values, that you would like to recibe and you recibe.
func Equal(got, want interface{}, t *testing.T) {
	c := pkg.ColorsPrint{}

	typeGot := reflect.TypeOf(got)
	typeWant := reflect.TypeOf(want)
	valueGot := reflect.ValueOf(got)
	valueWant := reflect.ValueOf(want)

	if typeGot != typeWant {
		c.Red(fmt.Sprintf("\n\tI recibe",typeGot))
		c.Red(fmt.Sprintf("\n\tBut I want",typeWant))
		t.Fail()
		return
	}
	if valueGot != valueWant {
		c.Red(fmt.Sprintf("\n\tI recibe",valueGot))
		c.Red(fmt.Sprintf("\n\tBut I want")
		fmt.Println(valueWant)
		t.Fail()
		return
	}

	c.Green("\n\tBoth values are equals")
}
