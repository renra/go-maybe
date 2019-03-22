package test

import (
  "fmt"
  "time"
  "app/maybe"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type TimeSuite struct {
  suite.Suite
}

func (s *TimeSuite) TestWithNilRef() {
  m := maybe.NewTime(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.DerefSafe()

  assert.Equal(s.T(), time.Time{}, value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Deref()
}

func (s *TimeSuite) TestWithValue() {
  input := time.Now()
  m := maybe.NewTime(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.DerefSafe()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Deref())
}

func (s *TimeSuite) TestValue() {
  input := time.Now()

  m := maybe.NewTime(&input)
  value, err := m.Value()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, value)

  m = maybe.NewTime(nil)
  value, err = m.Value()

  assert.Nil(s.T(), err)
  assert.Nil(s.T(), value)
}

func (s *TimeSuite) TestScan() {
  m := maybe.NewTime(nil)

  err := m.Scan(nil)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  // Time input
  input := time.Now()
  inputUnix := input.Unix()

  err = m.Scan(inputUnix)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), inputUnix, m.Deref().Unix())

  // String input
  err = m.Scan(fmt.Sprintf("%v", input))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}
