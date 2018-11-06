package nulls

import (
  "bytes"
  "database/sql"
  "encoding/json"
)

type String struct { sql.NullString }

func NewString(s string) (String) {
  return String{sql.NullString{String: s, Valid: true}}
}

func NewNullString() (String) {
  return String{sql.NullString{String: "", Valid: false}}
}

func (ns String) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return nullJSON, nil
	}
	return json.Marshal(ns.String)
}

func (ns *String) UnmarshalJSON(b []byte) error {
  var err error = nil
  if bytes.Equal(nullJSON, b) {
    ns.String = ""
    ns.Valid = false
  } else {
  	err = json.Unmarshal(b, &ns.String)
  	ns.Valid = (err == nil)
  }
	return err
}

func (ns String) IsEmpty() (bool) {
  return !ns.Valid || ns.String == ""
}

func (ns String) IsValid() (bool) {
  return ns.Valid
}
