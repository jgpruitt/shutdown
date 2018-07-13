// MIT License
//
// Copyright (c) 2018 John Pruitt
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package shutdown

import (
	"fmt"
	"testing"
)

func TestNow(t *testing.T) {
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
	t.Run("the first shutdown", func(t *testing.T) {
		if str != "cbaexit0" {
			t.Errorf("expected 'cbaexit0' but got '%s'", str)
		}
	})

	Now(6) // this one shouldn't do anything
	t.Run("the second shutdown", func(t *testing.T) {
		if str != "cbaexit0" {
			t.Errorf("expected 'cbaexit0' but got '%s'", str)
		}
	})
}
