package stringx_test

import (
	"backend/stringx"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestReverse(t *testing.T) {
	scenarios := []struct {
		input    string
		expected string
	}{
		{
			input:    "Hello",
			expected: "olleH",
		},
	}

	for i, scenario := range scenarios {
		testName := fmt.Sprintf("[%v] %v", i, scenario.input)
		t.Run(testName, func(t *testing.T) {
			assertEqual(t, scenario.expected, stringx.Reverse(scenario.input))
		})
	}
}

func assertEqual(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tact: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
