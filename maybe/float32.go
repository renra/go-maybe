package maybe

import (
  "errors"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewFloat32(ref *float32) Float32 {
  return Float32{ref: ref}
}

type Float32 struct {
  ref *float32
}

func (m Float32) HasValue() bool {
  return !(m.ref == nil)
}

func (m Float32) DerefSafe() (float32, *errtrace.Error) {
  if !m.HasValue() {
    return 0, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Float32) Deref() float32 {
  value, err := m.DerefSafe()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner interface
func (m *Float32) Scan(value interface{}) error {
  if value == nil {
    *m = NewFloat32(nil)
    return nil
  }

  valueFloat32, ok := value.(float32)

  if ok == false {
    *m = NewFloat32(nil)
    return errors.New(ConversionError)
  }

  *m = NewFloat32(&valueFloat32)

  return nil
}

func (m Float32) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Deref(), nil
  } else {
    return nil, nil
  }
}

