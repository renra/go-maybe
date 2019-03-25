package test

import (
  "fmt"
  "time"
  "app/maybe"
  "encoding/json"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type TimeSuite struct {
  suite.Suite
}

func (s *TimeSuite) TestWithNilRef() {
  m := maybe.NewTime(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), time.Time{}, value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Get()
}

func (s *TimeSuite) TestWithValue() {
  input := time.Now()
  m := maybe.NewTime(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Get())
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
  assert.Equal(s.T(), inputUnix, m.Get().Unix())

  // String input
  err = m.Scan(fmt.Sprintf("%v", input))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *TimeSuite) TestMarshalJSON() {
  m := maybe.NewTime(nil)
  bytes, err := m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte("null"), bytes)

  input := int64(2000)
  time := time.Unix(input, 0)

  m = maybe.NewTime(&time)
  bytes, err = m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte(fmt.Sprintf("%d", input)), bytes)
}

func (s *TimeSuite) TestUnmarshalJSON() {
  m := maybe.NewTime(nil)
  err := m.UnmarshalJSON([]byte("null"))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  input := int64(12)
  time := time.Unix(input, 0)

  err = m.UnmarshalJSON([]byte(fmt.Sprintf("%d", input)))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), time.Unix(), m.Get().Unix())

  err = m.UnmarshalJSON([]byte("foo"))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *TimeSuite) TestMarshalAndUnmarshalCycle() {
  input := time.Now()

  payload := struct{
    Field maybe.Time
  }{
    Field: maybe.NewTime(&input),
  }

  serializedPayload, err := json.Marshal(payload)
  assert.Nil(s.T(), err)

  err = json.Unmarshal(serializedPayload, &payload)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, payload.Field.HasValue())
  assert.Equal(s.T(), input.Unix(), payload.Field.Get().Unix())
}
