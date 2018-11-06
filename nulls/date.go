package nulls

import (
  "bytes"
  "database/sql"
  "encoding/json"
  "fmt"
	"reflect"
  "regexp"
  "strconv"
	"time"
)

// Date is 'timezone-less' so we base if off string as all we care about is
// YYYY-MM-DD format.
type Date struct { sql.NullString }
var dateRegexp *regexp.Regexp = regexp.MustCompile(`((\d{4})[\.-](\d\d)[\.-](\d\d))`)

func validateDate(s string) (error) {
  matches := dateRegexp.FindStringSubmatch(s)
  // Any invalid date results in an error.
  if matches == nil {
    return fmt.Errorf("'%s' does not parse as a date.", s)
  }
  var year, month, day int64
  var err error
  if year, err = strconv.ParseInt(matches[2], 10, 32); err != nil {
    return fmt.Errorf("'%s' does not parse as a date.", s)
  }
  if month, err = strconv.ParseInt(matches[3], 10, 32); err != nil {
    return fmt.Errorf("'%s' does not parse as a date.", s)
  }
  if day, err = strconv.ParseInt(matches[4], 10, 32); err != nil {
    return fmt.Errorf("'%s' does not parse as a date.", s)
  }
  // We use this to test that the string hits a valid day
  testTime := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

  // 'Date' normalizes, but we don't.
  if int(year) != testTime.Year() || time.Month(month) != testTime.Month() || int(day) != testTime.Day() {
    return fmt.Errorf("'%s' specifies a non-existent date (e.g., '2000-10-32').", s)
  }
  // Pull out just the 'YYYY-MM-DD' part; 'nt.String' will come in from MySQL with 'T00:00:00Z' on the end.
  return nil
}

func NewDate(s string) (Date, error) {
  if err := validateDate(s); err == nil {
    return Date{sql.NullString{String: s, Valid: true}}, nil
  } else {
    return NewNullDate(), err
  }
}

func NewNullDate() (Date) {
  return Date{sql.NullString{String: "", Valid: false}}
}

func (nt *Date) Scan(value interface{}) error {
	if reflect.TypeOf(value) == nil {
		*nt = NewNullDate()
    return nil
	} else {
    var s sql.NullString
  	if err := s.Scan(value); err != nil {
  		return err
  	} else {
      *nt, err = NewDate(s.String)
      return err
    }
  }
}

func (nd Date) MarshalJSON() ([]byte, error) {
  if !nd.Valid {
    return nullJSON, nil
  }
  return json.Marshal(nd.String)
}

func (nd *Date) UnmarshalJSON(b []byte) error {
  var err error = nil
  if bytes.Equal(nullJSON, b) {
    nd.String = ""
    nd.Valid = false
  } else {
    err = json.Unmarshal(b, &nd.String)
    if err == nil {
      err = validateDate(nd.String)
    }
    nd.Valid = (err == nil)
  }
  return err
}

func (n Date) IsEmpty() (bool) {
  return !n.Valid
}

func (n Date) IsValid() (bool) {
  return n.Valid
}
