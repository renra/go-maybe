package maybe

import (
  "time"
  "errors"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewTime(ref *time.Time) Time {
  return Time{ref: ref}
}

type Time struct {
  ref *time.Time
}

func (m Time) HasValue() bool {
  return !(m.ref == nil)
}

func (m Time) DerefSafe() (time.Time, *errtrace.Error) {
  if !m.HasValue() {
    return time.Time{}, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Time) Deref() time.Time {
  value, err := m.DerefSafe()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner interface
func (m *Time) Scan(value interface{}) error {
  if value == nil {
    *m = NewTime(nil)
    return nil
  }

  unixTimestamp, ok := value.(int64)

  if ok == false {
    *m = NewTime(nil)
    return errors.New(ConversionError)
  }

  timestamp := time.Unix(unixTimestamp, 0)
  *m = NewTime(&timestamp)

  return nil
}

func (m Time) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Deref(), nil
  } else {
    return nil, nil
  }
}
