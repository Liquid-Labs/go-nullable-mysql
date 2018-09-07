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

func TestNewNullBool(t *testing.T) {
  assertBoolNull(t, NewNullBool(), "NewNullBool()")
}

func TestUnmarshalBool(t *testing.T) {
	var b Bool
  // true
  panicIfErr(json.Unmarshal(boolTrueJson, &b))
	assertBoolValue(t, b, true, `json ` + boolTrueString)
  // false
  panicIfErr(json.Unmarshal(boolFalseJson, &b))
  assertBoolValue(t, b, false, `json ` + boolFalseString)
  // null
  panicIfErr(json.Unmarshal(nullJson, &b))
	assertBoolNull(t, b, `json ` + nullString)
  // bad type, thugh valid JSON
  assertError(t, json.Unmarshal(intJson, &b))
	assertBoolNull(t, b, `json ` + intString)
  // invalid json
  assertJsonSyntaxError(t, json.Unmarshal(invalidJson, &b))
	assertBoolNull(t, b, `json ` + invalidJsonString)
}

func TestMarshalBool(t *testing.T) {
	data, err := json.Marshal(NewBool(true))
	panicIfErr(err)
  assertJson(t, boolTrueJson, data)

  data, err = json.Marshal(NewBool(false))
  panicIfErr(err)
  assertJson(t, boolFalseJson, data)

  data, err = json.Marshal(NewNullBool())
  panicIfErr(err)
  assertJson(t, nullJson, data)
}

func TestBoolScan(t *testing.T) {
	var b Bool
	panicIfErr(b.Scan(true))
	assertBoolValue(t, b, true, "scanned true")

  panicIfErr(b.Scan(false))
	assertBoolValue(t, b, false, "scanned false")

	panicIfErr(b.Scan(nil))
	assertBoolNull(t, b, "scanned nil")
}

func assertBoolValue(t *testing.T, b Bool, expected bool, from string) {
  t.Helper()
	if expected != b.Bool {
		t.Errorf("Unexpected result from %s: %v ≠ %v\n", from, expected, b.Bool)
	}
	if !b.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}

func assertBoolNull(t *testing.T, b Bool, from string) {
  t.Helper()
	if b.Valid {
		t.Error(from, "is valid, but should be invalid")
	}
}
