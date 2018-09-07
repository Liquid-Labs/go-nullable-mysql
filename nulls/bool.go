package nulls

import (
  "bytes"
  "database/sql"
  "encoding/json"
)

type Bool struct { sql.NullBool }

func NewBool(b bool) (Bool) {
  return Bool{sql.NullBool{b, true}}
}

func NewNullBool() (Bool) {
  return Bool{sql.NullBool{false, false}}
}

func (nb Bool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return nullJSON, nil
	}
	return json.Marshal(nb.Bool)
}

func (nb *Bool) UnmarshalJSON(b []byte) error {
  var err error = nil
  if bytes.Equal(nullJSON, b) {
    nb.Bool = false
    nb.Valid = false
  } else {
  	err = json.Unmarshal(b, &nb.Bool)
  	nb.Valid = (err == nil)
  }
	return err
}

func (n Bool) IsEmpty() (bool) {
  return !n.Valid
}

func (n Bool) IsValid() (bool) {
  return n.Valid
}
