package maybe

import (
  "errors"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewInt8(ref *int8) Int8 {
  return Int8{ref: ref}
}

type Int8 struct {
  ref *int8
}

func (m Int8) HasValue() bool {
  return !(m.ref == nil)
}

func (m Int8) SafeGet() (int8, *errtrace.Error) {
  if !m.HasValue() {
    return 0, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Int8) Get() int8 {
  value, err := m.SafeGet()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner interface
func (m *Int8) Scan(value interface{}) error {
  if value == nil {
    *m = NewInt8(nil)
    return nil
  }

  valueInt8, ok := value.(int8)

  if ok == false {
    *m = NewInt8(nil)
    return errors.New(ConversionError)
  }

  *m = NewInt8(&valueInt8)

  return nil
}

func (m Int8) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Get(), nil
  } else {
    return nil, nil
  }
}



