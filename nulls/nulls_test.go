package nulls

import (
  "bytes"
  "encoding/json"
  "testing"
)

var (
  nullString = `null`
  nullJson  = []byte(nullString)
  invalidJsonString = `":)`
  invalidJson = []byte(invalidJsonString)
)

func assertJson(t *testing.T, expected []byte, result []byte) {
  t.Helper()
  if (!bytes.Equal(expected, result)) {
    t.Errorf("expected JSON '%s', but got '%s'", expected, result)
  }
}

func assertError(t *testing.T, err error) {
  t.Helper()
  if (err == nil) {
    t.Error("expected error, but got none.")
  }
}

func assertJsonSyntaxError(t *testing.T, err error) {
  t.Helper()
  assertError(t, err)
  if _, ok := err.(*json.SyntaxError); !ok {
    t.Errorf("expected json.SyntaxError, not %T", err)
  }
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
