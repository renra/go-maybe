package maybe

import (
  "fmt"
  "time"
  "errors"
  "strconv"
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

func (m Time) SafeGet() (time.Time, *errtrace.Error) {
  if !m.HasValue() {
    return time.Time{}, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Time) Get() time.Time {
  value, err := m.SafeGet()

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
    return m.Get(), nil
  } else {
    return nil, nil
  }
}

func (m Time) MarshalJSON() ([]byte, error) {
  if m.HasValue() {
    return []byte(fmt.Sprintf("%d", m.Get().Unix())), nil
  } else {
    return []byte("null"), nil
  }
}

func (m *Time) UnmarshalJSON(input []byte) error {
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

  time := time.Unix(value, 0)

  m.ref = &time
  return nil
}
