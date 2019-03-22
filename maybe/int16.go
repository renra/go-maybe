package maybe

import (
  "errors"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewInt16(ref *int16) Int16 {
  return Int16{ref: ref}
}

type Int16 struct {
  ref *int16
}

func (m Int16) HasValue() bool {
  return !(m.ref == nil)
}

func (m Int16) SafeGet() (int16, *errtrace.Error) {
  if !m.HasValue() {
    return 0, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Int16) Get() int16 {
  value, err := m.SafeGet()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner interface
func (m *Int16) Scan(value interface{}) error {
  if value == nil {
    *m = NewInt16(nil)
    return nil
  }

  valueInt16, ok := value.(int16)

  if ok == false {
    *m = NewInt16(nil)
    return errors.New(ConversionError)
  }

  *m = NewInt16(&valueInt16)

  return nil
}

func (m Int16) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Get(), nil
  } else {
    return nil, nil
  }
}
