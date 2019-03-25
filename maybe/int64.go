package maybe

import (
  "fmt"
  "errors"
  "strconv"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewInt64(ref *int64) Int64 {
  return Int64{ref: ref}
}

type Int64 struct {
  ref *int64
}

func (m Int64) HasValue() bool {
  return !(m.ref == nil)
}

func (m Int64) SafeGet() (int64, *errtrace.Error) {
  if !m.HasValue() {
    return 0, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Int64) Get() int64 {
  value, err := m.SafeGet()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner interface
func (m *Int64) Scan(value interface{}) error {
  if value == nil {
    *m = NewInt64(nil)
    return nil
  }

  valueInt64, ok := value.(int64)

  if ok == false {
    *m = NewInt64(nil)
    return errors.New(ConversionError)
  }

  *m = NewInt64(&valueInt64)

  return nil
}

func (m Int64) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Get(), nil
  } else {
    return nil, nil
  }
}

func (m Int64) MarshalJSON() ([]byte, error) {
  if m.HasValue() {
    return []byte(fmt.Sprintf("%d", m.Get())), nil
  } else {
    return []byte("null"), nil
  }
}

func (m *Int64) UnmarshalJSON(input []byte) error {
  inputStr := string(input)

  if inputStr == "null" {
    m.ref = nil
    return nil
  }

  value, err := strconv.ParseInt(inputStr, 10, 64)

  if err != nil {
    m.ref = nil
    return err
  }

  m.ref = &value
  return nil
}
