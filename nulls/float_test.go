package nulls

import (
	"encoding/json"
  "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
  assert.NoError(t, json.Unmarshal(noDecimalFloatJson, &f))
  assertFloatValue(t, f, 1.0, `json ` + noDecimalFloatString)
  // float with decimal values
  assert.NoError(t, json.Unmarshal(decimalFloatJson, &f))
  assertFloatValue(t, f, 1.23, `json ` + decimalFloatString)
  // null
  assert.NoError(t, json.Unmarshal(nullJson, &f))
	assertInvalid(t, f, `json ` + nullString)
  // bad type, though valid JSON with quoted decimal value
  assert.Error(t, json.Unmarshal(quotedDecimalFloatJson, &f))
	assertInvalid(t, f, `json "` + decimalFloatString + `"`)
  // invalid json
	assert.IsType(t, new(json.SyntaxError), json.Unmarshal(invalidJson, &f))
	assertInvalid(t, f, `json ` + invalidJsonString)
}

func TestMarshalFloat64(t *testing.T) {
	data, err := json.Marshal(NewFloat64(1))
	assert.NoError(t, err)
  assert.Equal(t, noDecimalFloatJson, data)

  data, err = json.Marshal(NewFloat64(1.23))
  assert.NoError(t, err)
  assert.Equal(t, decimalFloatJson, data)

  data, err = json.Marshal(NewNullFloat64())
  assert.NoError(t, err)
  assert.Equal(t, nullJson, data)
}

func TestFloat64Scan(t *testing.T) {
	var f Float64
	if assert.NoError(t, f.Scan(1)) {
		assertFloatValue(t, f, 1.00, "scanned 1")
	}

  if assert.NoError(t, f.Scan(1.23)) {
		assertFloatValue(t, f, 1.23, "scanned 1.23")
	}

	if assert.NoError(t, f.Scan(nil)) {
		assertInvalid(t, f, "scanned nil")
	}
}

func assertFloatValue(t *testing.T, f Float64, expected float64, from string) {
  t.Helper()
	if expected != f.Float64 {
		t.Errorf("Unexpected result from %s: %.2f ≠ %.2f\n", from, expected, f.Float64)
	}
  assertValid(t, f, from)
}
