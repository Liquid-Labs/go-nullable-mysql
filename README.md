# go-nullable-mysql
Enhanced support of `NULL` values in JSON marshalling and unmarshalling for MySQL datatypes in Go. I.e., SQL `NULL` <-> JSON `null`. Also, add support for MySQL `DATE` types.

# Usage
```go
import (
...
  "database/sql"
  "time"

  "github.com/Liquid-Labs/go-nullable-mysql/nulls"
)

type Foo struct {
  ADate       nulls.Date      `json:"aDate"`
  ATimestamp  nulls.Timestamp `json:"aTime"`
}

func SaveData(foo *Foo) (error) {
  _, err := insertFooStmt.Exec(foo.ADate, foo.ATimestamp)
  ...
}

func ScanFoo(row sql.Row) (*Foo, error) {
  foo := Foo{}
  if err := row.Scan(&foo.ADate, &foo.ATimestamp); err != nil {
    return nil, err
   }
   return &foo, nil
}

func MakeFoo(dateString string, time time.timeStamp) (Foo) {
  return Foo{nulls.NewDate(dateString), nulls.NewTimestamp(time)}
}
```

Supported types are:
* `NullBool` - a MySQL `BOOL` (==`TINYINT(1)`) <-> golang `bool`
* `NullDate` - a MySQL `DATE` <-> golang `string` in 'YYYY-MM-DD' format
* `NullFloat64` - a MySQL `DOUBLE` (or `FLOAT` if you like) <-> golang `float64`
* `NullInt64` - a MySQL `INT` <-> golang `int64`
* `NullString` - a MySQL `VARCHAR`, `CHAR`, `TEXT`, etc. <-> golang `string`
* `NullTimestamp` - a MySQL `TIMESTAMP` <-> golang `time.Time`

# References

* This work was based on implementation by [Supid Ravel](https://medium.com/@rsudip90) as discusesd in [this Medium article](https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267).
* [guregu/null](https://github.com/guregu/null) was used as an exapmle to
kickstart the unit tests. [guregu/null](https://github.com/guregu/null) is a
good alternative with similar functionality and a slightly different feature
set.
