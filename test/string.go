package test

import (
  "app/maybe"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type StringSuite struct {
  suite.Suite
}

func (s *StringSuite) TestWithNilRef() {
  m := maybe.NewString(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), "", value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Get()
}

func (s *StringSuite) TestWithValue() {
  input := "It is only with the heart that you can see well."
  m := maybe.NewString(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Get())
}

func (s *StringSuite) TestValue() {
  input := "It is only with the heart that you can see well."
  m := maybe.NewString(&input)
  value, err := m.Value()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, value)

  m = maybe.NewString(nil)
  value, err = m.Value()

  assert.Nil(s.T(), err)
  assert.Nil(s.T(), value)
}

func (s *StringSuite) TestScan() {
  m := maybe.NewString(nil)

  err := m.Scan(nil)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  stringInput := "What is essential is invisible to the eyes."
  err = m.Scan(stringInput)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), stringInput, m.Get())

  input := 12
  err = m.Scan(input)

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}
