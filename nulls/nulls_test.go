package nulls

import (
  "reflect"
  "runtime"
  "testing"

  "github.com/stretchr/testify/assert"
)

var (
  nullString = `null`
  nullJson  = []byte(nullString)
  invalidJsonString = `":)`
  invalidJson = []byte(invalidJsonString)
)

func testNullConstructors(t *testing.T) {
  constructors := []interface{}{NewNullBool, NewNullDate, NewNullFloat64, NewNullInt64, NewNullString, NewNullTimestamp}
  for _, constructor := range constructors {
    assertInvalid(t, constructor.(func()(Nullable))(), runtime.FuncForPC(reflect.ValueOf(constructor).Pointer()).Name() + `()`)
  }
}

// Helpers

func assertValid(t *testing.T, n Nullable, from string) {
  t.Helper()
  assert.Truef(t, n.IsValid(), "%s is invalid, but should be valid.", from)
  assert.Falsef(t, n.IsEmpty(), "%s is empty, but should be non-empty.", from)
}

func assertInvalid(t *testing.T, n Nullable, from string) {
  t.Helper()
  assert.Falsef(t, n.IsValid(), "%s is valid, but should be non-invalid.", from)
  assert.Truef(t, n.IsEmpty(), "%s is non-empty, but should be empty.", from)
}
