package nulls

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
  boolTrueString = `true`
	boolTrueJson  = []byte(boolTrueString)
  boolFalseString = `false`
	boolFalseJson = []byte(boolFalseString)
)

func TestNewBool(t *testing.T) {
	assertBoolValue(t, NewBool(true), true, "NewBool(true)")
	assertBoolValue(t, NewBool(false), false, "NewBool(false)")
}

func TestUnmarshalBool(t *testing.T) {
	var b Bool
  // true
  assert.NoError(t, json.Unmarshal(boolTrueJson, &b))
	assertBoolValue(t, b, true, `json ` + boolTrueString)
  // false
  assert.NoError(t, json.Unmarshal(boolFalseJson, &b))
  assertBoolValue(t, b, false, `json ` + boolFalseString)
  // null
  assert.NoError(t, json.Unmarshal(nullJson, &b))
	assertInvalid(t, b, `json ` + nullString)
  // bad type, though valid JSON
  assert.Error(t, json.Unmarshal(positiveIntJson, &b))
	assertInvalid(t, b, `json ` + positiveIntString)
  // invalid json
  assert.IsType(t, new(json.SyntaxError), json.Unmarshal(invalidJson, &b))
	assertInvalid(t, b, `json ` + invalidJsonString)
}

func TestMarshalBool(t *testing.T) {
	data, err := json.Marshal(NewBool(true))
	assert.NoError(t, err)
	assert.Equal(t, boolTrueJson, data)

  data, err = json.Marshal(NewBool(false))
  assert.NoError(t, err)
  assert.Equal(t, boolFalseJson, data)

  data, err = json.Marshal(NewNullBool())
  assert.NoError(t, err)
  assert.Equal(t, nullJson, data)
}

func TestBoolScan(t *testing.T) {
	var b Bool
	if assert.NoError(t, b.Scan(true)) {
		assertBoolValue(t, b, true, "scanned true")
	}

	if assert.NoError(t, b.Scan(false)) {
		assertBoolValue(t, b, false, "scanned false")
	}

	if assert.NoError(t, b.Scan(nil)) {
		assertInvalid(t, b, "scanned nil")
	}
}

func assertBoolValue(t *testing.T, b Bool, expected bool, from string) {
  t.Helper()
	assert.Equalf(t, expected, b.Bool, "Unexpected result from %s: %v ≠ %v\n", from, expected, b.Bool)
	assertValid(t, b, from)
}
