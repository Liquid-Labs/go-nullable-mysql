package nulls

import (
  "bytes"
  "fmt"
  "strings"
  "time"

  "github.com/go-sql-driver/mysql"
)

var nullTime, _ = time.Parse(time.RFC3339, "0000-00-00T00:00:00Z00:00")

type Timestamp struct { mysql.NullTime }

func NewTimestamp(t time.Time) (Timestamp) {
  return Timestamp{mysql.NullTime{t, true}}
}

func NewNullTimestamp() (Timestamp) {
  return Timestamp{mysql.NullTime{nullTime, false}}
}

func (nt Timestamp) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return nullJSON, nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

func (nt *Timestamp) UnmarshalJSON(b []byte) error {
  if bytes.Equal(nullJSON, b) {
    nt.Time = nullTime
    nt.Valid = false
  } else {
  	s := string(b)
  	s = strings.Trim(s, "\"")

  	x, err := time.Parse(time.RFC3339, s)
  	if err != nil {
  		nt.Valid = false
  		return err
  	}

  	nt.Time = x
  	nt.Valid = true
  }
	return nil
}

func (n Timestamp) IsEmpty() (bool) {
  return !n.Valid
}

func (n Timestamp) IsValid() (bool) {
  return n.Valid
}
