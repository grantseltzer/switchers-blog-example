package switchers

import (
	"testing"
)

func TestSwitch(t *testing.T) {

	var (
		a int
		b int
		c = &b
	)

	x := SwitchFunction(a, b, c)

	if x != "c" {
		t.Error("wtf?")
	}
}
