package nulls

import (
	"encoding/json"
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
    assertFloatValue(t, f, testValue, fmt.Sprintf(`NewFloat64(%f)`, testValue))
  }
}

func TestNewNullFloat(t *testing.T) {
  f := NewNullFloat64()
  if f.Float64 != 0.0 {
    t.Errorf("Expected t.Float64 to be 0.0, but instead got %.2f\n", f.Float64)
  }
  assertInvalid(t, f, `NewNullFloat64`)
}

func TestUnmarshalFloat64(t *testing.T) {
	var f Float64

  // int as float
  assertNoError(t, json.Unmarshal(noDecimalFloatJson, &f))
  assertFloatValue(t, f, 1.0, `json ` + noDecimalFloatString)
  // float with decimal values
  assertNoError(t, json.Unmarshal(decimalFloatJson, &f))
  assertFloatValue(t, f, 1.23, `json ` + decimalFloatString)
  // null
  assertNoError(t, json.Unmarshal(nullJson, &f))
	assertInvalid(t, f, `json ` + nullString)
  // bad type, though valid JSON with quoted decimal value
  assertError(t, json.Unmarshal(quotedDecimalFloatJson, &f))
	assertInvalid(t, f, `json "` + decimalFloatString + `"`)
  // invalid json
  assertJsonSyntaxError(t, json.Unmarshal(invalidJson, &f))
	assertInvalid(t, f, `json ` + invalidJsonString)
}

func TestMarshalFloat64(t *testing.T) {
	data, err := json.Marshal(NewFloat64(1))
	assertNoError(t, err)
  assertJson(t, noDecimalFloatJson, data)

  data, err = json.Marshal(NewFloat64(1.23))
  assertNoError(t, err)
  assertJson(t, decimalFloatJson, data)

  data, err = json.Marshal(NewNullFloat64())
  assertNoError(t, err)
  assertJson(t, nullJson, data)
}

func TestFloat64Scan(t *testing.T) {
	var f Float64
	panicIfErr(f.Scan(1))
	assertFloatValue(t, f, 1.00, "scanned 1")

  panicIfErr(f.Scan(1.23))
	assertFloatValue(t, f, 1.23, "scanned 1.23")

	panicIfErr(f.Scan(nil))
	assertInvalid(t, f, "scanned nil")
}

func assertFloatValue(t *testing.T, f Float64, expected float64, from string) {
  t.Helper()
	if expected != f.Float64 {
		t.Errorf("Unexpected result from %s: %.2f ≠ %.2f\n", from, expected, f.Float64)
	}
  assertValid(t, f, from)
}
