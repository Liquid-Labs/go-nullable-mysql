package nulls
// Thanks to Supid Raval
// https://gist.github.com/rsudip90/45fad7d8959c58bcc91d464873b50013
// to get us started, though that code did not handle unmarshalling nulls
// correctly. ;)

import (
  "bytes"
  "database/sql"
  "encoding/json"
  "fmt"
	"reflect"
  "regexp"
  "strconv"
  "strings"
	"time"

  "github.com/go-sql-driver/mysql"
)

var nullJSON = []byte("null")
var nullTime, _ = time.Parse(time.RFC3339, "0000-00-00T00:00:00Z00:00")

type Nullable interface{
  IsEmpty() (bool)
  IsValid() (bool)
}

type Int64 struct { sql.NullInt64 }

func NewInt64(i int64) (Int64) {
  return Int64{sql.NullInt64{i, true}}
}

func NewNullInt64() (Int64) {
  return Int64{sql.NullInt64{0, false}}
}

func (ni *Int64) MarshalJSON() ([]byte, error) {
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

func (ni *Int64) Native() *sql.NullInt64 {
  return &sql.NullInt64{Int64: ni.Int64, Valid: ni.Valid}
}

func (n *Int64) IsEmpty() (bool) {
  return !n.Valid
}

func (n *Int64) IsValid() (bool) {
  return n.Valid
}

// END: Null64Int handlers

type Bool struct { sql.NullBool }

func NewBool(b bool) (Bool) {
  return Bool{sql.NullBool{b, true}}
}

func NewNullBool() (Bool) {
  return Bool{sql.NullBool{false, false}}
}

func (nb *Bool) MarshalJSON() ([]byte, error) {
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

func (nb *Bool) Native() *sql.NullBool {
  return &sql.NullBool{Bool: nb.Bool, Valid: nb.Valid}
}

func (n *Bool) IsEmpty() (bool) {
  return !n.Valid
}

func (n *Bool) IsValid() (bool) {
  return n.Valid
}
// END NullBool handlers

type Float64 struct { sql.NullFloat64 }

func NewFloat64(f float64) (Float64) {
  return Float64{sql.NullFloat64{f, true}}
}

func NewNullFloat64() (Float64) {
  return Float64{sql.NullFloat64{0.0, false}}
}

func (nf *Float64) MarshalJSON() ([]byte, error) {
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

func (nf *Float64) Native() *sql.NullFloat64 {
  return &sql.NullFloat64{Float64: nf.Float64, Valid: nf.Valid}
}

func (n *Float64) IsEmpty() (bool) {
  return !n.Valid
}

func (n *Float64) IsValid() (bool) {
  return n.Valid
}
// END NullFloat64 handlers

// NullString is an alias for sql.NullString data type
type String struct { sql.NullString }

func NewString(s string) (String) {
  return String{sql.NullString{s, true}}
}

func NewNullString() (String) {
  return String{sql.NullString{"", false}}
}

func (ns *String) MarshalJSON() ([]byte, error) {
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

func (ns *String) Native() *sql.NullString {
  return &sql.NullString{String: ns.String, Valid: ns.Valid}
}

func (n *String) IsEmpty() (bool) {
  return !n.Valid || n.String == ""
}

func (n *String) IsValid() (bool) {
  return n.Valid
}
// END NullString handlers

type Timestamp struct { mysql.NullTime }

func NewTimestamp(t time.Time) (Timestamp) {
  return Timestamp{mysql.NullTime{t, true}}
}

func NewNullTimestamp() (Timestamp) {
  return Timestamp{mysql.NullTime{nullTime, false}}
}

func (nt *Timestamp) MarshalJSON() ([]byte, error) {
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

func (nt *Timestamp) Native() *mysql.NullTime {
  return &mysql.NullTime{Time: nt.Time, Valid: nt.Valid}
}

func (n *Timestamp) IsEmpty() (bool) {
  return !n.Valid
}

func (n *Timestamp) IsValid() (bool) {
  return n.Valid
}
// END NullTime handlers

// Date is 'timezone-less' so we base if off string as all we care about is
// YYYY-MM-DD format.
type Date struct { sql.NullString }
var dateRegexp *regexp.Regexp = regexp.MustCompile(`((\d{4})[\.-](\d\d)[\.-](\d\d))`)

func NewDate(s string) (Date) {
  return Date{sql.NullString{s, true}}
}

func NewNullDate() (Date) {
  return Date{sql.NullString{"", false}}
}

func (nt *Date) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*nt = NewNullDate()
	} else {
    matches := dateRegexp.FindStringSubmatch(s.String)
    // Any invalid date results in an error.
    if matches == nil {
      return fmt.Errorf("'%s' does not parse as a date.", s.String)
    }
    var year, month, day int64
    var err error
    if year, err = strconv.ParseInt(matches[2], 10, 32); err != nil {
      return fmt.Errorf("'%s' does not parse as a date.", s.String)
    }
    if month, err = strconv.ParseInt(matches[3], 10, 32); err != nil {
      return fmt.Errorf("'%s' does not parse as a date.", s.String)
    }
    if day, err = strconv.ParseInt(matches[4], 10, 32); err != nil {
      return fmt.Errorf("'%s' does not parse as a date.", s.String)
    }
    // We use this to test that the string hits a valid day
    testTime := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

    // 'Date' normalizes, but we don't.
    if int(year) != testTime.Year() || time.Month(month) != testTime.Month() || int(day) != testTime.Day() {
      return fmt.Errorf("'%s' specifies a non-existent date (e.g., '2000-10-32').", s.String)
    }
    // Pull out just the 'YYYY-MM-DD' part; 'nt.String' will come in from MySQL with 'T00:00:00Z' on the end.
		*nt = NewDate(matches[1])
	}

	return nil
}

func (nd *Date) MarshalJSON() ([]byte, error) {
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
    nd.Valid = (err == nil)
  }
  return err
}

func (nd *Date) Native() *sql.NullString {
  return &sql.NullString{String: nd.String, Valid: nd.Valid}
}

func (n *Date) IsEmpty() (bool) {
  return !n.Valid
}

func (n *Date) IsValid() (bool) {
  return n.Valid
}
// END NullDate handlers
