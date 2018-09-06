package nulls

var nullJSON = []byte("null")

type Nullable interface{
  IsEmpty() (bool)
  IsValid() (bool)
}
