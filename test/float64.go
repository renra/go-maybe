package test

import (
  "fmt"
  "app/maybe"
  "encoding/json"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type Float64Suite struct {
  suite.Suite
}

func (s *Float64Suite) TestWithNilRef() {
  m := maybe.NewFloat64(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), float64(0), value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Get()
}

func (s *Float64Suite) TestWithValue() {
  input := float64(9)
  m := maybe.NewFloat64(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Get())
}

func (s *Float64Suite) TestValue() {
  input := float64(12)

  m := maybe.NewFloat64(&input)
  value, err := m.Value()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, value)

  m = maybe.NewFloat64(nil)
  value, err = m.Value()

  assert.Nil(s.T(), err)
  assert.Nil(s.T(), value)
}

func (s *Float64Suite) TestScan() {
  m := maybe.NewFloat64(nil)

  err := m.Scan(nil)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  // Float64 input
  input := float64(12)
  err = m.Scan(input)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Get())

  // String input
  err = m.Scan(fmt.Sprintf("%f", input))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *Float64Suite) TestMarshalJSON() {
  m := maybe.NewFloat64(nil)
  bytes, err := m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte("null"), bytes)

  input := float64(12.5)
  m = maybe.NewFloat64(&input)
  bytes, err = m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte(fmt.Sprintf("%.1f", input)), bytes)
}

func (s *Float64Suite) TestUnmarshalJSON() {
  m := maybe.NewFloat64(nil)
  err := m.UnmarshalJSON([]byte("null"))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  input := float64(12.5)
  err = m.UnmarshalJSON([]byte(fmt.Sprintf("%f", input)))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Get())

  err = m.UnmarshalJSON([]byte("foo"))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *Float64Suite) TestMarshalAndUnmarshalCycle() {
  input := 12.5

  payload := struct{
    Field maybe.Float64
  }{
    Field: maybe.NewFloat64(&input),
  }

  serializedPayload, err := json.Marshal(payload)
  assert.Nil(s.T(), err)

  err = json.Unmarshal(serializedPayload, &payload)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, payload.Field.HasValue())
  assert.Equal(s.T(), input, payload.Field.Get())
}
