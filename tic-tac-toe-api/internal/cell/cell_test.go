package cell

import (
	"testing"
)

func TestValidateReturnsFalseForAnyInvalidValueXOE(t *testing.T) {
	const maxByte = ^Value(0)

	for i := E + 1; i < maxByte; i++ {
		if validate(i) != false {
			t.Errorf("invalid cell value was validated")
		}
	}
}

func TestNextValueSwitchesBetweenXAndO(t *testing.T) {
	var data = [...]struct {
		in  Value
		out Value
	}{
		{O, X},
		{X, O},
	}

	for i := 0; i < len(data); i++ {
		next := nextValue(data[i].in)

		if next != data[i].out {
			t.Errorf("wrong calculation of the NextValue for %v, (%v) given; (%v) expected", data[i].in, next, data[i].out)
		}
	}
}