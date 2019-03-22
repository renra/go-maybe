package test

import (
  "fmt"
  "app/maybe"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type Int32Suite struct {
  suite.Suite
}

func (s *Int32Suite) TestWithNilRef() {
  m := maybe.NewInt32(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.DerefSafe()

  assert.Equal(s.T(), int32(0), value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Deref()
}

func (s *Int32Suite) TestWithValue() {
  input := int32(9)
  m := maybe.NewInt32(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.DerefSafe()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Deref())
}

func (s *Int32Suite) TestValue() {
  input := int32(12)

  m := maybe.NewInt32(&input)
  value, err := m.Value()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, value)

  m = maybe.NewInt32(nil)
  value, err = m.Value()

  assert.Nil(s.T(), err)
  assert.Nil(s.T(), value)
}

func (s *Int32Suite) TestScan() {
  m := maybe.NewInt32(nil)

  err := m.Scan(nil)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  // Int32 input
  input := int32(12)
  err = m.Scan(input)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Deref())

  // String input
  err = m.Scan(fmt.Sprintf("%d", input))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}
