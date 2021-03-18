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

	c.Cyan("Assert Equal " + description)

	valueGot := reflect.ValueOf(got)
	valueWant := reflect.ValueOf(want)

	if valueWant != valueGot {
		c.Red(fmt.Sprint("\n\tI recibe", got))
		c.Red(fmt.Sprint("\n\tBut I want", want))
		t.Fail()
		return
	}

	c.Green(fmt.Sprint("\n\t I recibe ", want))
}

func IsNil(got interface{}, description string, t *testing.T) {
	c := pkg.ColorsPrint{}

	c.Cyan("Assert Nil " + description)

	if !reflect.ValueOf(got).IsNil() {
		c.Red(fmt.Sprint("\n\tI recibe", got))
		t.Fail()
		return
	}

	c.Green("\n\t I recibe a nil value")
}

func IsNotNil(got interface{}, description string, t *testing.T) {
	c := pkg.ColorsPrint{}

	c.Cyan("Assert Nil " + description)

	if reflect.ValueOf(got).IsNil() {
		c.Red(fmt.Sprint("\n\tI recibe nil value"))
		t.Fail()
		return
	}

	c.Green(fmt.Sprint("\n\tI recibe", got))
}
