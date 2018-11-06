package nulls

import (
  "bytes"
  "database/sql"
  "encoding/json"
)

type Int64 struct { sql.NullInt64 }

func NewInt64(i int64) (Int64) {
  return Int64{sql.NullInt64{Int64: i, Valid: true}}
}

func NewNullInt64() (Int64) {
  return Int64{sql.NullInt64{Int64: 0, Valid: false}}
}

func (ni Int64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return nullJSON, nil
	}
	return json.Marshal(ni.Int64)
}

func (ni *Int64) UnmarshalJSON(b []byte) error {
  var err error = nil
  if bytes.Equal(nullJSON, b) {
    ni.Int64 = 0
    ni.Valid = false
  } else {
  	err = json.Unmarshal(b, &ni.Int64)
  	ni.Valid = (err == nil)
  }
	return err
}

func (n Int64) IsEmpty() (bool) {
  return !n.Valid
}

func (n Int64) IsValid() (bool) {
  return n.Valid
}
