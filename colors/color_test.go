package color

import (
	"fmt"
	"testing"
)

type testValue struct {
	actual interface{}
	expect string
}

func TestPackage(t *testing.T) {
	testValues := []testValue{
		testValue{Black("text"), "\033[30mtext\033[0m"},
		testValue{WhiteBg(fmt.Sprintf("%s", Black("text"))), "\033[47m\033[30mtext\033[0m\033[0m"},
	}
	for _, value := range testValues {
		if fmt.Sprintf("%s", value.actual) != value.expect {
			t.Logf("%#v != %#v", fmt.Sprintf("%s", value.actual), value.expect)
			t.Fail()
		}
	}
}
