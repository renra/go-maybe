package maybe

import (
  "errors"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewInt(ref *int) Int {
  return Int{ref: ref}
}

type Int struct {
  ref *int
}

func (m Int) HasValue() bool {
  return !(m.ref == nil)
}

func (m Int) SafeGet() (int, *errtrace.Error) {
  if !m.HasValue() {
    return 0, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Int) Get() int {
  value, err := m.SafeGet()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner interface
func (m *Int) Scan(value interface{}) error {
  if value == nil {
    *m = NewInt(nil)
    return nil
  }

  valueInt, ok := value.(int)

  if ok == false {
    *m = NewInt(nil)
    return errors.New(ConversionError)
  }

  *m = NewInt(&valueInt)

  return nil
}

func (m Int) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Get(), nil
  } else {
    return nil, nil
  }
}
