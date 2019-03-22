package test

import (
  "testing"
  "github.com/stretchr/testify/suite"
)

func TestInt(t *testing.T) {
  suite.Run(t, new(IntSuite))
}

func TestInt64(t *testing.T) {
  suite.Run(t, new(Int64Suite))
}

func TestInt32(t *testing.T) {
  suite.Run(t, new(Int32Suite))
}

func TestInt16(t *testing.T) {
  suite.Run(t, new(Int16Suite))
}

func TestInt8(t *testing.T) {
  suite.Run(t, new(Int8Suite))
}

func TestFloat64(t *testing.T) {
  suite.Run(t, new(Float64Suite))
}

func TestFloat32(t *testing.T) {
  suite.Run(t, new(Float32Suite))
}

func TestString(t *testing.T) {
  suite.Run(t, new(StringSuite))
}
