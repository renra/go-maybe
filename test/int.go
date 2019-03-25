package test

import (
  "fmt"
  "app/maybe"
  "encoding/json"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type IntSuite struct {
  suite.Suite
}

func (s *IntSuite) TestWithNilRef() {
  m := maybe.NewInt(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), 0, value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Get()
}

func (s *IntSuite) TestWithValue() {
  input := 9
  m := maybe.NewInt(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Get())
}

func (s *IntSuite) TestValue() {
  input := 12

  m := maybe.NewInt(&input)
  value, err := m.Value()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, value)

  m = maybe.NewInt(nil)
  value, err = m.Value()

  assert.Nil(s.T(), err)
  assert.Nil(s.T(), value)
}

func (s *IntSuite) TestScan() {
  m := maybe.NewInt(nil)

  err := m.Scan(nil)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  // Int input
  input := 12
  err = m.Scan(input)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Get())

  // String input
  err = m.Scan(fmt.Sprintf("%d", input))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *IntSuite) TestMarshalJSON() {
  m := maybe.NewInt(nil)
  bytes, err := m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte("null"), bytes)

  input := 12
  m = maybe.NewInt(&input)
  bytes, err = m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte(fmt.Sprintf("%d", input)), bytes)
}

func (s *IntSuite) TestUnmarshalJSON() {
  m := maybe.NewInt(nil)
  err := m.UnmarshalJSON([]byte("null"))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  input := 12
  err = m.UnmarshalJSON([]byte(fmt.Sprintf("%d", input)))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Get())

  err = m.UnmarshalJSON([]byte("foo"))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *IntSuite) TestMarshalAndUnmarshalCycle() {
  input := 12

  payload := struct{
    Field maybe.Int
  }{
    Field: maybe.NewInt(&input),
  }

  serializedPayload, err := json.Marshal(payload)
  assert.Nil(s.T(), err)

  err = json.Unmarshal(serializedPayload, &payload)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, payload.Field.HasValue())
  assert.Equal(s.T(), input, payload.Field.Get())
}
