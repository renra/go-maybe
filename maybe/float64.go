package maybe

import (
  "errors"
  "strconv"
  "encoding/json"
  "database/sql/driver"
  "github.com/renra/go-errtrace/errtrace"
)

func NewFloat64(ref *float64) Float64 {
  return Float64{ref: ref}
}

type Float64 struct {
  ref *float64
}

func (m Float64) HasValue() bool {
  return !(m.ref == nil)
}

func (m Float64) SafeGet() (float64, *errtrace.Error) {
  if !m.HasValue() {
    return 0, errtrace.New(DereferenceError)
  }

  return *m.ref, nil
}

// Convenient but unsafe. Use at your own risk after checking HasValue()
func (m Float64) Get() float64 {
  value, err := m.SafeGet()

  if err != nil {
    panic(err)
  }

  return value
}

// Returns a normal error, otherwise does not fulfill the Scanner interface
func (m *Float64) Scan(value interface{}) error {
  if value == nil {
    *m = NewFloat64(nil)
    return nil
  }

  valueFloat64, ok := value.(float64)

  if ok == false {
    *m = NewFloat64(nil)
    return errors.New(ConversionError)
  }

  *m = NewFloat64(&valueFloat64)

  return nil
}

func (m Float64) Value() (driver.Value, error) {
  if m.HasValue() {
    return m.Get(), nil
  } else {
    return nil, nil
  }
}

func (m Float64) MarshalJSON() ([]byte, error) {
  if m.HasValue() {
    return json.Marshal(m.Get())
  } else {
    return json.Marshal(nil)
  }
}

func (m *Float64) UnmarshalJSON(input []byte) error {
  inputStr := string(input)

  if inputStr == "null" {
    m.ref = nil
    return nil
  }

  value, err := strconv.ParseFloat(inputStr, 64)

  if err != nil {
    m.ref = nil
    return err
  }

  m.ref = &value
  return nil
}

