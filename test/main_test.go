package test

import (
  "testing"
  "github.com/stretchr/testify/suite"
)

func TestInt(t *testing.T) {
  suite.Run(t, new(IntSuite))
}