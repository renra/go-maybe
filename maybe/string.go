package maybe

import (
  "errors"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewString(ref *string) String {
  return String{ref: ref}
}

type String struct {
  ref *string
}

func (m String) HasValue() bool {
  return !(m.ref == nil)
}

func (m String) SafeGet() (string, *errtrace.Error) {
  if !m.HasValue() {
    return "", errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m String) Get() string {
  value, err := m.SafeGet()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner stringerface
func (m *String) Scan(value interface{}) error {
  if value == nil {
    *m = NewString(nil)
    return nil
  }

  valueString, ok := value.(string)

  if ok == false {
    *m = NewString(nil)
    return errors.New(ConversionError)
  }

  *m = NewString(&valueString)

  return nil
}

func (m String) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Get(), nil
  } else {
    return nil, nil
  }
}

