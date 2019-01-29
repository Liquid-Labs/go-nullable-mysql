package nulls

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
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
  assert.NoError(t, err)
	assertDateValue(t, d, validDateString, "NewDate(" + validDateString + ")")

  d, err = NewDate(invalidMonthString)
  assert.Error(t, err)
	assertInvalid(t, d, "NewDate(" + invalidMonthString + ")")

  d, err = NewDate(invalidDayString)
  assert.Error(t, err)
  assertInvalid(t, d, "NewDate(" + invalidDayString + ")")

  d, err = NewDate(positiveIntString)
  assert.Error(t, err)
  assertInvalid(t, d, "NewDate(" + positiveIntString + ")")

	d, err = NewDate(nonIntYearString)
	assert.Error(t, err)
	assertInvalid(t, d, "NewDate(" + nonIntYearString + ")")
}

func TestUnmarshalDate(t *testing.T) {
	testStrings := []string{validDateString, invalidMonthString, invalidDayString, positiveIntString}
  testJsons := [][]byte{validDateJson, invalidMonthJson, invalidDayJson, positiveIntJson}
  const lastGood = 0
  var d Date
  for i, dateJson := range testJsons {
    if i <= lastGood {
      assert.NoError(t, json.Unmarshal(dateJson, &d))
    	assertDateValue(t, d, testStrings[i], `json ` + string(dateJson))
    } else {
      assert.Error(t, json.Unmarshal(dateJson, &d))
      assertInvalid(t, d, `json ` + string(dateJson))
    }
  }

	assert.NoError(t, json.Unmarshal(nullJson, &d))
	assertInvalid(t, d, `json ` + nullString)
}

func TestMarshalDate(t *testing.T) {
  d, err := NewDate(validDateString)
  // note, we already tested the 'err' above, so we're not gonig to do it again
  data, err := json.Marshal(d)
  assert.NoError(t, err)
  assert.Equal(t, validDateJson, data)
}

func TestDateScan(t *testing.T) {
	var d Date
	if assert.NoError(t, d.Scan(validDateString)) {
		assertDateValue(t, d, validDateString, "scanned " + validDateString)
	}
}

func assertDateValue(t *testing.T, d Date, expected string, from string) {
  t.Helper()
	if expected != d.String {
		t.Errorf("Unexpected result from %s: %v ≠ %v\n", from, expected, d.String)
	}
  assertValid(t, d, from)
}
