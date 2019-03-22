package test

import (
  "fmt"
  "app/maybe"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type Int16Suite struct {
  suite.Suite
}

func (s *Int16Suite) TestWithNilRef() {
  m := maybe.NewInt16(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.DerefSafe()

  assert.Equal(s.T(), int16(0), value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Deref()
}

func (s *Int16Suite) TestWithValue() {
  input := int16(9)
  m := maybe.NewInt16(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.DerefSafe()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Deref())
}

func (s *Int16Suite) TestValue() {
  input := int16(12)

  m := maybe.NewInt16(&input)
  value, err := m.Value()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, value)

  m = maybe.NewInt16(nil)
  value, err = m.Value()

  assert.Nil(s.T(), err)
  assert.Nil(s.T(), value)
}

func (s *Int16Suite) TestScan() {
  m := maybe.NewInt16(nil)

  err := m.Scan(nil)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  // Int16 input
  input := int16(12)
  err = m.Scan(input)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Deref())

  // String input
  err = m.Scan(fmt.Sprintf("%d", input))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}


