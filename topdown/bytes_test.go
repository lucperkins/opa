package topdown

import (
	"github.com/open-policy-agent/opa/ast"
	"testing"
)

func runNumBytesTest(t *testing.T, note, rule string, expected int) {
	s := ast.StringTerm(rule).Value
	val, err := builtinNumBytes(s)
	if err != nil {
		t.Fatalf("numbytes: %v", err)
	}

	i := val.(ast.Number)

	num, ok := i.Int()
	if !ok {
		t.Fatalf("numbytes failed num parse")
	}

	if num != expected {
		t.Fatalf(`numbytes failure on "%s": expected value %d does not match %d`, note, expected, num)
	}
}

func TestNumBytes(t *testing.T) {
	t.Run("Passing", func(t *testing.T) {
		tests := []struct {
			note     string
			rule     string
			expected int
		}{
			{"zero", `0`, 0},
			{"raw number", `12345`, 12345},
			{"10 kilobytes uppercase", `10KB`, 10 * 1024},
			{"10 kilobytes lowercase", `10kb`, 10 * 1024},
			{"200 megabytes as m", `200m`, 200 * 1024 * 1024},
			{"200 megabytes as mb", `200mb`, 200 * 1024 * 1024},
			{"37 gigabytes", `37g`, 37 * 1024 * 1024 * 1024},
			{"3 terabytes", `3t`, 3 * 1024 * 1024 * 1024 * 1024},
		}

		for _, tc := range tests {
			runNumBytesTest(t, tc.note, tc.rule, tc.expected)
		}
	})
}
