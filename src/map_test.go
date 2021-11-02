package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T){
    assert.Equal(t, Sum(1,4), 5)  
}
