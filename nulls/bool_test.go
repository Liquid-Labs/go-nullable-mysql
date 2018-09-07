package nulls

import (
	"encoding/json"
	"testing"
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
  assertNoError(t, json.Unmarshal(boolTrueJson, &b))
	assertBoolValue(t, b, true, `json ` + boolTrueString)
  // false
  assertNoError(t, json.Unmarshal(boolFalseJson, &b))
  assertBoolValue(t, b, false, `json ` + boolFalseString)
  // null
  assertNoError(t, json.Unmarshal(nullJson, &b))
	assertInvalid(t, b, `json ` + nullString)
  // bad type, though valid JSON
  assertError(t, json.Unmarshal(positiveIntJson, &b))
	assertInvalid(t, b, `json ` + positiveIntString)
  // invalid json
  assertJsonSyntaxError(t, json.Unmarshal(invalidJson, &b))
	assertInvalid(t, b, `json ` + invalidJsonString)
}

func TestMarshalBool(t *testing.T) {
	data, err := json.Marshal(NewBool(true))
	assertNoError(t, err)
  assertJson(t, boolTrueJson, data)

  data, err = json.Marshal(NewBool(false))
  assertNoError(t, err)
  assertJson(t, boolFalseJson, data)

  data, err = json.Marshal(NewNullBool())
  assertNoError(t, err)
  assertJson(t, nullJson, data)
}

func TestBoolScan(t *testing.T) {
	var b Bool
	panicIfErr(b.Scan(true))
	assertBoolValue(t, b, true, "scanned true")

  panicIfErr(b.Scan(false))
	assertBoolValue(t, b, false, "scanned false")

	panicIfErr(b.Scan(nil))
	assertInvalid(t, b, "scanned nil")
}

func assertBoolValue(t *testing.T, b Bool, expected bool, from string) {
  t.Helper()
	if expected != b.Bool {
		t.Errorf("Unexpected result from %s: %v ≠ %v\n", from, expected, b.Bool)
	}
	assertValid(t, b, from)
}
