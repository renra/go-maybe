package test

import (
  "fmt"
  "app/maybe"
  "encoding/json"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type Int8Suite struct {
  suite.Suite
}

func (s *Int8Suite) TestWithNilRef() {
  m := maybe.NewInt8(nil)

  assert.Equal(s.T(), false, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), int8(0), value)
  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), maybe.DereferenceError, err.Error())

  defer func() {
    r := recover()

    assert.NotNil(s.T(), r)
    assert.Equal(s.T(), maybe.DereferenceError, err.Error())
  }()

  m.Get()
}

func (s *Int8Suite) TestWithValue() {
  input := int8(9)
  m := maybe.NewInt8(&input)

  assert.Equal(s.T(), true, m.HasValue())

  value, err := m.SafeGet()

  assert.Equal(s.T(), input, value)
  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, m.Get())
}

func (s *Int8Suite) TestValue() {
  input := int8(12)

  m := maybe.NewInt8(&input)
  value, err := m.Value()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), input, value)

  m = maybe.NewInt8(nil)
  value, err = m.Value()

  assert.Nil(s.T(), err)
  assert.Nil(s.T(), value)
}

func (s *Int8Suite) TestScan() {
  m := maybe.NewInt8(nil)

  err := m.Scan(nil)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  // Int8 input
  input := int8(12)
  err = m.Scan(input)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Get())

  // String input
  err = m.Scan(fmt.Sprintf("%d", input))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *Int8Suite) TestMarshalJSON() {
  m := maybe.NewInt8(nil)
  bytes, err := m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte("null"), bytes)

  input := int8(12)
  m = maybe.NewInt8(&input)
  bytes, err = m.MarshalJSON()

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), []byte(fmt.Sprintf("%d", input)), bytes)
}

func (s *Int8Suite) TestUnmarshalJSON() {
  m := maybe.NewInt8(nil)
  err := m.UnmarshalJSON([]byte("null"))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())

  input := int8(12)
  err = m.UnmarshalJSON([]byte(fmt.Sprintf("%d", input)))

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, m.HasValue())
  assert.Equal(s.T(), input, m.Get())

  err = m.UnmarshalJSON([]byte("foo"))

  assert.NotNil(s.T(), err)
  assert.Equal(s.T(), false, m.HasValue())
}

func (s *Int8Suite) TestMarshalAndUnmarshalCycle() {
  input := int8(12)

  payload := struct{
    Field maybe.Int8
  }{
    Field: maybe.NewInt8(&input),
  }

  serializedPayload, err := json.Marshal(payload)
  assert.Nil(s.T(), err)

  err = json.Unmarshal(serializedPayload, &payload)

  assert.Nil(s.T(), err)
  assert.Equal(s.T(), true, payload.Field.HasValue())
  assert.Equal(s.T(), input, payload.Field.Get())
}
