package nulls

import (
	// "encoding/json"
  "fmt"
	"testing"
)

var (
  noDecimalFloatString = `1`
	noDecimalFloatJson  = []byte(noDecimalFloatString)
  decimalFloatString = `1.23`
	decimalFloatJson = []byte(decimalFloatString)
  quotedDecimalFloatJson = []byte(`"` + decimalFloatString + `"`)
)

func TestNewFloat(t *testing.T) {
  testValues := []float64{0, -1, 1.23, -1.23}
  for _, testValue := range testValues {
    f := NewFloat64(testValue)
    assertFloatValue(t, f, testValue, fmt.Sprintf(`new Float64(%f)`, testValue))
  }
}

// TODO: Unmarshal tests

// TODO: Marshal tests

// TODO: Scan tests

func assertFloatValue(t *testing.T, f Float64, expected float64, from string) {
  t.Helper()
	if expected != f.Float64 {
		t.Errorf("Unexpected result from %s: %.2f ≠ %.2f\n", from, expected, f.Float64)
	}
  assertValid(t, f, from)
}
