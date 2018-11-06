package nulls

import (
  "bytes"
  "database/sql"
  "encoding/json"
)

type Float64 struct { sql.NullFloat64 }

func NewFloat64(f float64) (Float64) {
  return Float64{sql.NullFloat64{Float64: f, Valid: true}}
}

func NewNullFloat64() (Float64) {
  return Float64{sql.NullFloat64{Float64: 0.0, Valid: false}}
}

func (nf Float64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return nullJSON, nil
	}
	return json.Marshal(nf.Float64)
}

func (nf *Float64) UnmarshalJSON(b []byte) error {
  var err error = nil
  if bytes.Equal(nullJSON, b) {
    nf.Float64 = 0.0
    nf.Valid = false
  } else {
  	err = json.Unmarshal(b, &nf.Float64)
  	nf.Valid = (err == nil)
  }
	return err
}

func (n Float64) IsEmpty() (bool) {
  return !n.Valid
}

func (n Float64) IsValid() (bool) {
  return n.Valid
}
