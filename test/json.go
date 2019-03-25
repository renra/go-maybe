package test

import (
  "fmt"
  "time"
  "app/maybe"
  "encoding/json"
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
)

type JSONSuite struct {
  suite.Suite
}

type payload struct {
  I maybe.Int            `json:,omitempty`
  I8 maybe.Int8
  I16 maybe.Int16
  I32 maybe.Int32
  I64 maybe.Int64
  F32 maybe.Float32
  F64 maybe.Float64
  S maybe.String
  T maybe.Time
}

func (s *JSONSuite) TestMarshalingAndUnmarshaling() {
  p := payload{}

  output, err := json.Marshal(p)
  assert.Nil(s.T(), err)

  inputString := "Some string"

  //inputInt := 12
  inputInt8 := int8(13)
  inputInt16 := int16(14)
  inputInt32 := int32(15)
  inputInt64 := int64(16)

  inputFloat32 := float32(1.4)
  inputFloat64 := float64(3.14)

  inputTime := time.Now()

  p = payload{
    S: maybe.NewString(&inputString),
    //I: maybe.NewInt(&inputInt),
    I8: maybe.NewInt8(&inputInt8),
    I16: maybe.NewInt16(&inputInt16),
    I32: maybe.NewInt32(&inputInt32),
    I64: maybe.NewInt64(&inputInt64),
    F32: maybe.NewFloat32(&inputFloat32),
    F64: maybe.NewFloat64(&inputFloat64),
    T: maybe.NewTime(&inputTime),
  }

  output, err = json.Marshal(p)

  fmt.Println(fmt.Sprintf("%v", string(output)))
  fmt.Println(fmt.Sprintf("%v", p))

  output, err = json.Marshal(p)
  assert.Nil(s.T(), err)

  //fmt.Println(fmt.Sprintf("%v", err.Error()))
}
