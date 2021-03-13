package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/KummerCompany/kairos/pkg"
)

// Equal : Compare two values, that you would like to recibe and you recibe.
func Equal(got, want interface{}, description string, t *testing.T) {
	c := pkg.ColorsPrint{}

	c.Cyan("Assert Eqaul" + description)

	typeGot := reflect.TypeOf(got)
	typeWant := reflect.TypeOf(want)
	valueGot := reflect.ValueOf(got)
	valueWant := reflect.ValueOf(want)

	if typeGot != typeWant {
		c.Red(fmt.Sprintf("\n\tI recibe", typeGot))
		c.Red(fmt.Sprintf("\n\tBut I want", typeWant))
		t.Fail()
		return
	}
	if valueGot != valueWant {
		c.Red(fmt.Sprintf("\n\tI recibe", valueGot))
		c.Red(fmt.Sprintf("\n\tBut I want", valueWant))
		t.Fail()
		return
	}

	c.Green("\n\t I recibe " + valueGot.String())
}
