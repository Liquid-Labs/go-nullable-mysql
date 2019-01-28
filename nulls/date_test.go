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
	nonIntYearString = `abc-05-08`
	nonIntYearJson = []byte(`"` + nonIntYearString + `"`)
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

  d, err = NewDate(positiveIntString)
  assertError(t, err)
  assertInvalid(t, d, "NewDate(" + positiveIntString + ")")

	d, err = NewDate(nonIntYearString)
	assertError(t, err)
	assertInvalid(t, d, "NewDate(" + nonIntYearString + ")")
}

func TestUnmarshalDate(t *testing.T) {
	testStrings := []string{validDateString, invalidMonthString, invalidDayString, positiveIntString}
  testJsons := [][]byte{validDateJson, invalidMonthJson, invalidDayJson, positiveIntJson}
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

	assertNoError(t, json.Unmarshal(nullJson, &d))
	assertInvalid(t, d, `json ` + nullString)
}

func TestMarshalDate(t *testing.T) {
  d, err := NewDate(validDateString)
  // note, we already tested the 'err' above, so we're not gonig to do it again
  data, err := json.Marshal(d)
  assertNoError(t, err)
  assertJson(t, validDateJson, data)
}

func TestDateScan(t *testing.T) {
	var d Date
	panicIfErr(d.Scan(validDateString))
	assertDateValue(t, d, validDateString, "scanned " + validDateString)
}

func assertDateValue(t *testing.T, d Date, expected string, from string) {
  t.Helper()
	if expected != d.String {
		t.Errorf("Unexpected result from %s: %v ≠ %v\n", from, expected, d.String)
	}
  assertValid(t, d, from)
}
