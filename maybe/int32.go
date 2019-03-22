package maybe

import (
  "errors"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewInt32(ref *int32) Int32 {
  return Int32{ref: ref}
}

type Int32 struct {
  ref *int32
}

func (m Int32) HasValue() bool {
  return !(m.ref == nil)
}

func (m Int32) DerefSafe() (int32, *errtrace.Error) {
  if !m.HasValue() {
    return 0, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Int32) Deref() int32 {
  value, err := m.DerefSafe()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner interface
func (m *Int32) Scan(value interface{}) error {
  if value == nil {
    *m = NewInt32(nil)
    return nil
  }

  valueInt32, ok := value.(int32)

  if ok == false {
    *m = NewInt32(nil)
    return errors.New(ConversionError)
  }

  *m = NewInt32(&valueInt32)

  return nil
}

func (m Int32) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Deref(), nil
  } else {
    return nil, nil
  }
}
