package test

import (
  "fmt"
  "app/maybe"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type Float32Suite struct {
  suite.Suite
}

func (s *Float32Suite) TestWithNilRef() {
  m := maybe.NewFloat32(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), float32(0), value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Get()
}

func (s *Float32Suite) TestWithValue() {
  input := float32(9)
  m := maybe.NewFloat32(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Get())
}

func (s *Float32Suite) TestValue() {
  input := float32(12)

  m := maybe.NewFloat32(&input)
  value, err := m.Value()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, value)

  m = maybe.NewFloat32(nil)
  value, err = m.Value()

  assert.Nil(s.T(), err)
  assert.Nil(s.T(), value)
}

func (s *Float32Suite) TestScan() {
  m := maybe.NewFloat32(nil)

  err := m.Scan(nil)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  // Float32 input
  input := float32(12)
  err = m.Scan(input)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Get())

  // String input
  err = m.Scan(fmt.Sprintf("%f", input))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *Float32Suite) TestMarshalJSON() {
  m := maybe.NewFloat32(nil)
  bytes, err := m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte("null"), bytes)

  input := float32(12.5)
  m = maybe.NewFloat32(&input)
  bytes, err = m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte(fmt.Sprintf("%.1f", input)), bytes)
}

func (s *Float32Suite) TestUnmarshalJSON() {
  m := maybe.NewFloat32(nil)
  err := m.UnmarshalJSON([]byte("null"))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  input := float32(12.5)
  err = m.UnmarshalJSON([]byte(fmt.Sprintf("%f", input)))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Get())

  err = m.UnmarshalJSON([]byte("foo"))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}
