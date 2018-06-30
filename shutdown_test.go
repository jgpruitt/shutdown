package shutdown

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {
	var str = ""

	exit = func(code int) {
		// don't actually exit the process for the purposes of testing
		str = fmt.Sprintf("%sexit%d", str, code)
	}
	block = func() {
		// don't actually block forever for the purposes of testing
	}

	Add(func() {
		str = str + "a"
	})
	Add(func() {
		str = str + "b"
	})
	Add(func() {
		str = str + "c"
	})

	go func() {
		OnSignal()
	}()

	Now(0)
	Now(6) // this one shouldn't do anything

	if str != "cbaexit0" {
		t.Errorf("expected 'cbaexit0' but got '%s'", str)
	}
}
