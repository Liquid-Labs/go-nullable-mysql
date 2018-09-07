package nulls

import (
	"encoding/json"
	"testing"
)

var (
  validDateString = `2018-05-08`
	validDateJson  = []byte(`"` + validDateString + `"`)
  invalidMonthString = `2018-13-08`
	invalidMonthJson = []byte(`"` + invalidMonthString + `"`)
  invalidDayString = `2018-06-31`
  invalidDayJson = []byte(`"` + invalidDayString + `"`)
)

func TestNewDate(t *testing.T) {
  d, err := NewDate(validDateString)
  assertNoError(t, err)
	assertDateValue(t, d, validDateString, "NewDate(" + validDateString + ")")

  d, err = NewDate(invalidMonthString)
  assertError(t, err)
	assertInvalid(t, d, "NewDate(" + invalidMonthString + ")")

  d, err = NewDate(invalidDayString)
  assertError(t, err)
  assertInvalid(t, d, "NewDate(" + invalidDayString + ")")

  d, err = NewDate(intString)
  assertError(t, err)
  assertInvalid(t, d, "NewDate(" + intString + ")")
}

func TestNewNullDate(t *testing.T) {
  assertInvalid(t, NewNullBool(), "NewNullDate()")
}

func TestUnmarshalDate(t *testing.T) {
	testStrings := []string{validDateString, invalidMonthString, invalidDayString, intString}
  testJsons := [][]byte{validDateJson, invalidMonthJson, invalidDayJson, intJson}
  const lastGood = 0
  var d Date
  for i, dateJson := range testJsons {
    if i <= lastGood {
      assertNoError(t, json.Unmarshal(dateJson, &d))
    	assertDateValue(t, d, testStrings[i], `json ` + string(dateJson))
    } else {
      assertError(t, json.Unmarshal(dateJson, &d))
      assertInvalid(t, d, `json ` + string(dateJson))
    }
  }
}
/*
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
*/
func assertDateValue(t *testing.T, d Date, expected string, from string) {
  t.Helper()
	if expected != d.String {
		t.Errorf("Unexpected result from %s: %v ≠ %v\n", from, expected, d.String)
	}
	if !d.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}
